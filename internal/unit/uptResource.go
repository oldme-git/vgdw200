package unit

import (
	"github.com/oldme-git/vgdw200/internal/message"
)

type UptResourceReady struct {
}

func (u UptResourceReady) Sender(data []byte) []byte {
	return message.New().Cmd(0x81).Data(data).Get()
}

func (u UptResourceReady) Recver(in *message.VgDw200In) ([]byte, error) {
	return nil, nil
}

type UptResourceTransmit struct {
}

func (u UptResourceTransmit) Sender(data []byte) []byte {
	return message.New().Cmd(0x82).Data(data).Get()
}

func (u UptResourceTransmit) Recver(in *message.VgDw200In) ([]byte, error) {
	return nil, nil
}

type UptResourceOver struct {
}

func (u UptResourceOver) Sender(data []byte) []byte {
	return message.New().Cmd(0x83).Data(data).Get()
}

func (u UptResourceOver) Recver(in *message.VgDw200In) ([]byte, error) {
	return nil, nil
}
