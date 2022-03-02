package ziface

// 不同的路由提供不同指令，可以把不同消息处理不同方式

type IRouter interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
