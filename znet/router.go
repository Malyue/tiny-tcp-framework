package znet

import "zinx/ziface"

// 创建基类，实现Router时，根据需求对基类方法进行重写
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req ziface.IRequest)  {}
func (br *BaseRouter) Handle(req ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(req ziface.IRequest) {}
