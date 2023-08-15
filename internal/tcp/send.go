package tcp

import (
	"fmt"
	"github.com/gogf/gf/v2/net/gtcp"
)

func Send(conn *gtcp.Conn, data []byte) error {
	fmt.Printf("send：%X\n", data)
	err := conn.Send(data)
	if err != nil {
		return err
	}
	return nil
}
