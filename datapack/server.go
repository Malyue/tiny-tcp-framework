package main

import (
	"fmt"
	"io"
	"net"
	"zinx/znet"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("server listen err : ", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
		}
		go func(conn net.Conn) {
			// 创建封包拆包对象
			dp := znet.NewDataPack()

			for {
				headData := make([]byte, dp.GetHeadLen())
				// ReadFull会把headData填充满为止
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head error")
					break
				}
				// 将headData字节流拆包到msg中
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack err : ", err)
					break
				}
				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					// 根据dataLen从io中读取字节流
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack data err : ", err)
						return
					}
					fmt.Println("==> Recv Msg:ID= ", msg.Id, ",len=", msg.Len, ",data=", string(msg.Data))
				}
			}
		}(conn)
	}
}
