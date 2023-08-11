package unit

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oldme-git/vgdw200/internal/message"
)

type ShowMsg struct {
}

func (s ShowMsg) Sender(data []byte) []byte {
	return message.New().Cmd(0x61).Data(data).Get()
}

func (s ShowMsg) Recver(in *message.VgDw200In) ([]byte, error) {
	j, _ := json.Marshal(map[string]string{
		"ack": "ok",
		"msg": "I'm Ok, What about you?",
	})
	return s.Sender(j), nil
}

type ShowImg struct {
}

func (s ShowImg) Sender(data []byte) []byte {
	return message.New().Cmd(0x63).Data(data).Get()
}

func (s ShowImg) Recver(in *message.VgDw200In) ([]byte, error) {
	if in.Flag != 0 {
		return nil, errors.New(fmt.Sprintf("更新图片失败, 标识字：%d", in.Flag))
	}
	return nil, nil
}

type InWindow struct {
}

func (i InWindow) Sender(data []byte) []byte {
	return message.New().Cmd(0x64).Data(data).Get()
}

func (i InWindow) Recver(in *message.VgDw200In) ([]byte, error) {
	if in.Flag != 0 {
		return nil, errors.New(fmt.Sprintf("进入窗口失败, 标识字：%d", in.Flag))
	}
	return nil, nil
}
