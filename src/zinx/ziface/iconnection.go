package ziface

import "net"

type Iconnection interface {

	// Start 启动连接
	Start()
	// Stop 结束当前链接工作
	Stop()
	// GetTCPConnection 获取当前绑定的TCP stocked
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前模块链接的ID
	GetConnID() uint32
	// RemoteAddr 获取客户端的TCP状态 IP Port
	RemoteAddr() net.Addr
	// Send 发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
