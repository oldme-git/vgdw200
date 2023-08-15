package unit

import (
	"errors"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/tcp"
)

type FindDeviceStatus struct {
}

func (f FindDeviceStatus) Sender(conn *gtcp.Conn, data []byte) error {
	return nil
}

func (f FindDeviceStatus) Recver(conn *gtcp.Conn, data *message.VgDw200In) error {
	if data.Flag == 0x00 {
		return nil
	}
	return errors.New("设备状态异常")
}

type BuzzerAndLED struct {
}

func (b BuzzerAndLED) Sender(conn *gtcp.Conn, data []byte) error {
	return nil
}

func (b BuzzerAndLED) Recver(conn *gtcp.Conn, data *message.VgDw200In) error {
	return nil
}

type AudioCtrl struct {
}

func (a AudioCtrl) Sender(conn *gtcp.Conn, data []byte) error {
	msg := message.New().Cmd(0x29).Data(data).Get()
	err := tcp.Send(conn, msg)
	if err != nil {
		return err
	}
	return nil
}

func (a AudioCtrl) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	if in.Flag != 0 {
		return newErr(0x29, in.Flag)
	}
	return nil
}
