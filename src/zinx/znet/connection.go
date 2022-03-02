package znet

import (
	"Game_Zinx/src/zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	// 当前链接的套接字
	Conn *net.TCPConn
	// 连接ID
	ConnID uint32
	// 连接状态
	isClosed bool

	// 告知当前连接已经退出
	ExitChan chan bool
	// 该链接处理的方法Router
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID = ", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recev buf err", err)
			continue
		}
		// d得到当前Conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		//从路由Routers 中找到注册绑定Conn的对应Handle
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		// 调用路由，从Conn中找到注册绑定对应router
	}

}

// Start 启动连接
func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	go c.StartReader()

}

// Stop 结束当前链接工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	// 回收资源
	close(c.ExitChan)

}

// GetTCPConnection 获取当前绑定的TCP stocked
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前模块链接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取客户端的TCP状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
