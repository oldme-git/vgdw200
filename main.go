package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/oldme-git/vgdw200/internal/logic"
	"github.com/oldme-git/vgdw200/internal/message"
	"github.com/oldme-git/vgdw200/internal/service"
	_ "github.com/oldme-git/vgdw200/internal/service"
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

func (t *TcpList) broadcast(msg []byte) {
	defer t.Unlock()
	t.Lock()
	for k, v := range t.list {
		err := v.Send(msg)
		if err != nil {
			fmt.Printf("发送信息失败了,addr:%s, err: %s", k, err.Error())
		}
	}
}

func main() {
	tcpList := NewTcpList()
	go func() {
		s := g.Server()
		s.BindHandler("/", func(r *ghttp.Request) {
			//cmd := r.Get("msg").String()
			srv, _ := service.SRV.Get(0x81)
			msg := srv.Sender([]byte{0x02})
			//msg := message.New().Cmd(0x63).Data([]byte{0x02}).Get()
			fmt.Printf("send1：%X\n", msg)
			tcpList.broadcast(msg)
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
				err := logic.Recver(data, func(in *message.VgDw200In, back []byte) {
					if in.Cmd == 0x81 {
						// 传输包的处理
						err := logic.UptResource(conn, "./1.zip")
						if err != nil {
							g.Log().Error(gctx.New(), err)
						}
						return
					}
					// 默认处理行为
					if len(back) > 0 {
						if err := conn.Send(back); err != nil {
							g.Log().Error(gctx.New(), err)
						}
					}
				})
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
