package unit

import (
	"crypto/md5"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/oldme-git/vgdw200/internal/logic"
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/tcp"
	"time"
)

type UptResourceReady struct {
}

func (u UptResourceReady) Sender(conn *gtcp.Conn, data []byte) error {
	resource, err := logic.NewResource(string(data))
	if err != nil {
		return err
	}
	var (
		d        = make([]byte, 7)
		size     = uint32(resource.GetSize())
		sizeByte = message.DataBytesBig32(size)
		num      = uint16(resource.GetPkgNum())
		numByte  = message.DataBytesBig16(num)
	)
	d[0] = 0x02
	copy(d[1:3], numByte[0:2])
	copy(d[3:7], sizeByte[0:4])

	msg := message.New().Cmd(0x81).Data(d).Get()
	return tcp.Send(conn, msg)
}

func (u UptResourceReady) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	if in.Flag != 0x00 {
		return newErr(0x81, in.Flag)
	}
	unit, err := srv.Get(0x82)
	if err != nil {
		return err
	}
	// 发送传输指令
	err = unit.Sender(conn, []byte(FileName))
	if err != nil {
		return err
	}
	return nil
}

type UptResourceTransmit struct {
}

func (u UptResourceTransmit) Sender(conn *gtcp.Conn, data []byte) error {
	r, err := logic.NewResource(string(data))
	defer r.Close()
	if err != nil {
		return err
	}
	err = r.Transmit(func(i uint, bytes []byte) error {
		// 获取序号
		var (
			order = message.DataBytesBig16(uint16(i))
			data  = make([]byte, len(bytes)+2)
		)
		copy(data[0:2], order[:])
		copy(data[2:], bytes)
		msg := message.New().Cmd(0x82).Data(data).Get()
		// 缓一下发送，不然小盒子反应不过来
		time.Sleep(10 * time.Millisecond)
		err := tcp.Send(conn, msg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// 发送结束指令
	unit, err := srv.Get(0x83)
	if err != nil {
		return err
	}
	err = unit.Sender(conn, []byte(FileName))
	return nil
}

func (u UptResourceTransmit) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	return nil
}

type UptResourceOver struct {
}

func (u UptResourceOver) Sender(conn *gtcp.Conn, data []byte) error {
	r, err := logic.NewResource(string(data))
	defer r.Close()
	if err != nil {
		return err
	}
	d := make([]byte, 0)
	err = r.Transmit(func(i uint, bytes []byte) error {
		d = append(d, bytes...)
		return nil
	})
	if err != nil {
		return err
	}
	hasher := md5.New()
	hasher.Write(d)
	md5Byte := hasher.Sum(nil)

	d1 := make([]byte, 18)
	copy(d1[2:18], md5Byte[:])
	msg := message.New().Cmd(0x83).Data(d1).Get()
	err = tcp.Send(conn, msg)
	if err != nil {
		return err
	}
	return nil
}

func (u UptResourceOver) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	if in.Flag != 0 {
		return newErr(0x83, in.Flag)
	}
	return nil
}
