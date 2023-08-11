package unit

import (
	"errors"
	"github.com/oldme-git/vgdw200/internal/message"
)

type FindDeviceStatus struct {
}

func (f FindDeviceStatus) Sender(data []byte) []byte {
	return message.New().Cmd(0x01).Get()
}

func (f FindDeviceStatus) Recver(in *message.VgDw200In) ([]byte, error) {
	if in.Flag == 0x00 {
		return nil, nil
	}
	return nil, errors.New("设备状态异常")
}

type BuzzerAndLED struct {
}

func (b BuzzerAndLED) Sender(data []byte) []byte {
	return message.New().Cmd(0x04).Data(data).Get()
}

func (b BuzzerAndLED) Recver(in *message.VgDw200In) ([]byte, error) {
	return nil, nil
}

type AudioCtrl struct {
}

func (a AudioCtrl) Sender(data []byte) []byte {
	return message.New().Cmd(0x29).Data(data).Get()
}

func (a AudioCtrl) Recver(in *message.VgDw200In) ([]byte, error) {
	return nil, nil
}
