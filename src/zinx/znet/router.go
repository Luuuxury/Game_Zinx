package znet

import "Game_Zinx/src/zinx/ziface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request ziface.IRouter)  {}
func (br *BaseRouter) Handle(request ziface.IRouter)     {}
func (br *BaseRouter) PostHandle(request ziface.IRouter) {}
