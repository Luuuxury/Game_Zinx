package znet

import (
	"Game_Zinx/src/zinx/utils"
	"Game_Zinx/src/zinx/ziface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	fmt.Printf("[Start] Server Listenner at IP:%s, Port %d, is startinh \n", s.IP, s.Port)
	// 1. 获取一个TCP Addr （套接字）
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}
	// 2.监听地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IPVersion, "err", err)
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
	//先初始化全局配置文件
	utils.GlobalObject.Reload()

	s := &Server{
		Name:      utils.GlobalObject.Name, //从全局参数获取
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,    //从全局参数获取
		Port:      utils.GlobalObject.TcpPort, //从全局参数获取
		Router:    nil,
	}
	return s
}
