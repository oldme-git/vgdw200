package service

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/oldme-git/vgdw200/internal/message"
)

type SrvInterface interface {
	// Sender 发送信息
	Sender(conn *gtcp.Conn, data []byte) error
	// Recver 接收信息
	Recver(conn *gtcp.Conn, data *message.VgDw200In) error
}

type Service struct {
	instances map[byte]SrvInterface
	service   map[byte]func() SrvInterface
}

func NewService() *Service {
	return &Service{
		instances: make(map[byte]SrvInterface),
		service:   make(map[byte]func() SrvInterface),
	}
}

func (s *Service) Register(name byte, srv func() SrvInterface) {
	s.service[name] = srv
}

func (s *Service) Get(name byte) (SrvInterface, error) {
	if srv, ok := s.instances[name]; ok {
		return srv, nil
	}
	if srv, ok := s.service[name]; ok {
		s.instances[name] = srv()
		return s.instances[name], nil
	}

	return nil, gerror.Newf("服务未注册 0x%x", name)
}
