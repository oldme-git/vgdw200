package logic

import (
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/service"
)

func Recver(data []byte, f func(*message.VgDw200In, []byte)) error {
	in := message.Parser(data)
	srv, err := service.SRV.Get(in.Cmd)
	if err != nil {
		return err
	}
	recver, err := srv.Recver(in)
	if err != nil {
		return err
	}
	f(in, recver)
	return nil
}
