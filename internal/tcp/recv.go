package tcp

import (
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/service"
)

var srv *service.Service

func Init(s *service.Service) {
	srv = s
}

func Recv(conn *gtcp.Conn, data []byte) error {
	in := message.Parser(data)
	unit, err := srv.Get(in.Cmd)
	if err != nil {
		return err
	}
	err = unit.Recver(conn, in)
	if err != nil {
		return err
	}
	return nil
}
