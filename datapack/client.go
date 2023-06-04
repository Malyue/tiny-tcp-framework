package main

import (
	"fmt"
	"net"
	"zinx/znet"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client dial err : ", err)
		return
	}

	// 创建一个封包对象
	dp := znet.NewDataPack()
	// 封装msg包
	msg := &znet.Message{
		Id:   0,
		Len:  5,
		Data: []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData, err := dp.Pack(msg)
	if err != nil {
		return
	}

	msg1 := &znet.Message{
		Id:   1,
		Len:  7,
		Data: []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		return
	}

	sendData = append(sendData, sendData1...)
	conn.Write(sendData)
	select {}
}
