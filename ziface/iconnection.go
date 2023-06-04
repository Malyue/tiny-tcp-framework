package ziface

import "net"

// 连接封装和业务绑定

type IConnection interface {
	// 启动连接，让当前连接开始工作
	Start()
	// 停止连接，结束当前连接状态M
	Stop()
	// 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接ID
	GetConnID() uint32
	// 获取远程客户端地址信息
	RemoteAddr() net.Addr
	// 发送Message数据,进行封包
	SendMsg(msgId uint32, data []byte) error
}

type HandFunc func(*net.TCPConn, []byte, int) error
