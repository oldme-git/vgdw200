package unit

import (
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/tcp"
)

type ShowMsg struct {
}

func (s ShowMsg) Sender(conn *gtcp.Conn, data []byte) error {
	return nil
}

func (s ShowMsg) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	//j, _ := json.Marshal(map[string]string{
	//	"ack": "ok",
	//	"msg": "I'm Ok, What about you?",
	//})
	return nil
}

type ShowImg struct {
}

func (s ShowImg) Sender(conn *gtcp.Conn, data []byte) error {
	msg := message.New().Cmd(0x63).Data(data).Get()
	err := tcp.Send(conn, msg)
	if err != nil {
		return err
	}
	return nil
}

func (s ShowImg) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	if in.Flag != 0 {
		return newErr(0x63, in.Flag)
	}
	return nil
}

type InWindow struct {
}

func (i InWindow) Sender(conn *gtcp.Conn, data []byte) error {
	msg := message.New().Cmd(0x64).Data(data).Get()
	err := tcp.Send(conn, msg)
	if err != nil {
		return err
	}
	return nil
}

func (i InWindow) Recver(conn *gtcp.Conn, in *message.VgDw200In) error {
	if in.Flag != 0 {
		return newErr(0x64, in.Flag)
	}
	return nil
}
