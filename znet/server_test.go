package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test ... start")
	time.Sleep(1 * time.Second)

	// connect server
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		_, err := conn.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("write conn err: ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		fmt.Println("server call back: ", string(buf[:cnt]), cnt)
		time.Sleep(1 * time.Second)
	}

}

func TestServer(t *testing.T) {
	s := NewServer("zinx v0.1")

	go ClientTest()
	s.Serve()
}
