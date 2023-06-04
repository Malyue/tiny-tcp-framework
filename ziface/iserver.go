package ziface

// 服务器接口，包括一个起点，暂停和开启业务服务的方法
type IServer interface {
	// method of start the server
	Start()
	// method of stop the server
	Stop()
	// start the server
	Serve()
	// 给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(router IRouter)
}

type HandleFunc interface {
}
