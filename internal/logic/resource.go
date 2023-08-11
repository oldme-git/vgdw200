package logic

import (
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/service"
	"io"
	"os"
)

type Resource struct {
	f *os.File
	// 分包大小，默认1024
	pkgNum   uint
	FileInfo os.FileInfo
}

// UptResourceReady 下发开始发送文件指令

// UptResource 分包传输文件
func UptResource(conn *gtcp.Conn, path string) error {
	var (
		err error
	)
	r, err := NewResource(path)
	defer r.Close()
	if err != nil {
		return err
	}
	serviceTransmit, err := service.SRV.Get(0x82)
	if err != nil {
		return err
	}
	var data []byte
	err = r.Transmit(func(i uint, bytes []byte) error {
		// 获取数据长度，大端序模式
		l := message.DataBytesBig(uint16(len(bytes)))
		// 拼接数据
		data = append(l[0:], bytes...)
		serviceTransmit.Sender(data)
		err = conn.Send(data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func NewResource(path string) (*Resource, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return &Resource{
		f:        f,
		pkgNum:   1024,
		FileInfo: fileInfo,
	}, nil
}

// Transmit 将包拆解，结果请注入依赖处理
func (r *Resource) Transmit(f func(uint, []byte) error) error {
	var (
		b = make([]byte, 1024)
		i uint
	)

	for {
		_, err := r.f.Read(b)
		if err == io.EOF {
			break
		}
		err = f(i, b)
		if err != nil {
			return err
		}
		i++
	}
	return nil
}

// GetSize 获取文件大小
func (r *Resource) GetSize() uint {
	return uint(r.FileInfo.Size())
}

// GetPkgNum 获取分包数量
func (r *Resource) GetPkgNum() (num uint) {
	num = r.GetSize() / r.pkgNum
	if r.GetSize()%r.pkgNum > 0 {
		num++
	}
	return
}

func (r *Resource) Close() error {
	err := r.f.Close()
	if err != nil {
		return err
	}
	return nil
}
