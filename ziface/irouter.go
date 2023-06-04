package ziface

/*
	路由接口
*/

type IRouter interface {
	// 处理conn业务之前的钩子方法
	PreHandle(request IRequest)
	// 处理conn业务的方法
	Handle(reqeust IRequest)
	// 处理conn业务之后的钩子方法
	PostHandle(request IRequest)
}
