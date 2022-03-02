package znet

import (
	"Game_Zinx/src/zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int
	Router    ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port %d, is startinh \n", s.Ip, s.Port)
	// 1. 获取一个TCP Addr （套接字）
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}
	// 2.监听地址
	listenner, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IpVersion, "err", err)
		return
	}
	fmt.Println("Start Zinx server succ,", s.Name, "succ Listenning...")

	var cid uint32
	cid = 0
	// 3.阻塞等待客户端连接, 处理客户端业务 （读写）
	for {
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}
		// 将处理新连接的业务方法 和 conn 进行绑定， 得到我们的链接模块
		dealConn := NewConnection(conn, cid, s.Router)
		cid++
		// 启动当前的链接业务处理
		go dealConn.Start()

	}

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// TODO 做一些启动之后额外的业务
	select {}
}
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router succ...")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
