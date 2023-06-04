package ziface

// 抽象IRquest层，把连接和请求信息包装进来

type IRequest interface {
	// 获取请求连接信息
	GetConnection() IConnection
	// 获取请求消息的数据
	GetData() []byte
}
