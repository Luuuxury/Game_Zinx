package znet

import (
	"Game_Zinx/src/zinx/ziface"
	"fmt"
	"net"
	"time"
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
	fmt.Println("[STOP] Zinx server , name ", s.Name)
	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理

}

func (s *Server) Serve() {
	s.Start()
	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
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
