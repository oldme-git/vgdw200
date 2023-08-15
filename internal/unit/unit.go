package unit

import "github.com/oldme-git/vgdw200/internal/service"

var srv *service.Service

func Init(s *service.Service) {
	srv = s
}
