package logic

import (
	"io"
	"os"
)

type Resource struct {
	f *os.File
	// 分包大小，默认1024
	pkgNum   uint
	FileInfo os.FileInfo
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
