package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/oldme-git/vgdw200/internal/service"
	"github.com/oldme-git/vgdw200/internal/tcp"
	"github.com/oldme-git/vgdw200/internal/unit"
	"sync"
)

type TcpList struct {
	list map[string]*gtcp.Conn
	sync.RWMutex
}

func NewTcpList() *TcpList {
	return &TcpList{
		list: make(map[string]*gtcp.Conn),
	}
}

func (t *TcpList) addTcp(addr string, conn *gtcp.Conn) {
	defer t.Unlock()
	t.Lock()
	t.list[addr] = conn
}

func (t *TcpList) delTcp(addr string) {
	defer t.Unlock()
	t.Lock()
	delete(t.list, addr)
}

func (t *TcpList) broadcast(unit service.SrvInterface, msg []byte) {
	defer t.Unlock()
	t.Lock()
	var err error
	for k, v := range t.list {
		err = unit.Sender(v, msg)
		if err != nil {
			fmt.Printf("发送信息失败了,addr:%s, err: %s", k, err.Error())
		}
	}
}

func main() {
	srv := service.NewService()
	registerSrv(srv)

	tcpList := NewTcpList()
	go func() {
		s := g.Server()
		s.BindHandler("/", func(r *ghttp.Request) {
			//cmd := r.Get("msg").String()
			u, _ := srv.Get(0x81)
			tcpList.broadcast(u, []byte(unit.FileName))
			//u, _ := srv.Get(0x63)
			//tcpList.broadcast(u, []byte("bk"))
			//u, _ := srv.Get(0x29)
			//tcpList.broadcast(u, []byte{0x04})
		})
		s.SetPort(8888)
		s.Run()
	}()
	err := gtcp.NewServer("192.168.10.49:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		addr := conn.RemoteAddr().String()
		tcpList.addTcp(addr, conn)
		fmt.Printf("打开tcp连接%s \n", addr)

		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				err := tcp.Recv(conn, data)
				if err != nil {
					g.Log().Error(gctx.New(), err)
				}
			}
			if err != nil {
				fmt.Printf("tcp连接关闭%s \n", addr)
				tcpList.delTcp(addr)
				break
			}
		}
	}).Run()
	if err != nil {
		panic(err)
	}
}

// 注册单元
func registerSrv(srv *service.Service) {
	srv.Register(0x01, func() service.SrvInterface {
		return &unit.FindDeviceStatus{}
	})
	srv.Register(0x04, func() service.SrvInterface {
		return &unit.BuzzerAndLED{}
	})
	srv.Register(0x29, func() service.SrvInterface {
		return &unit.AudioCtrl{}
	})
	srv.Register(0x61, func() service.SrvInterface {
		return &unit.ShowMsg{}
	})
	srv.Register(0x63, func() service.SrvInterface {
		return &unit.ShowImg{}
	})
	srv.Register(0x64, func() service.SrvInterface {
		return &unit.InWindow{}
	})
	srv.Register(0x81, func() service.SrvInterface {
		return &unit.UptResourceReady{}
	})
	srv.Register(0x82, func() service.SrvInterface {
		return &unit.UptResourceTransmit{}
	})
	srv.Register(0x83, func() service.SrvInterface {
		return &unit.UptResourceOver{}
	})
	registerPkg(srv)
}

// 注册srv到包中
func registerPkg(srv *service.Service) {
	tcp.Init(srv)
	unit.Init(srv)
}
