package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/ziface"
	"zinx/zinx/utils"
)

// implement the interface of IServer
type Server struct {
	// server name
	Name string
	// tcp4 or other
	IPVersion string
	// server ip
	IP string
	// server port
	Port int
	// 当前Server由用户绑定的回调router，也就是Server注册的链接对应的处理业务
	Router ziface.IRouter
}

var _ ziface.IServer = &Server{}

func NewServer(name ...string) ziface.IServer {
	if len(name) > 1 {
		panic("the server name is too long,it is limit 1")
	}
	var serverName string
	config := utils.GetGlobalConfig()
	if len(name) == 1 {
		serverName = name[0]
	} else {
		serverName = config.Name
	}
	s := &Server{
		Name:      serverName,
		IPVersion: "tcp4",
		IP:        config.Host,
		Port:      config.TcpPort,
		Router:    nil,
	}
	return s
}

// 实现Server服务器启动方法
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP : %s, PORT : %d ,is starting\n", s.IP, s.Port)

	// start a goroutine to listen
	// 开始监听服务器地址和端口
	go func() {
		// returns an addr of ip endpoint
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve ip addr error: ", err)
			return
		}

		// listen the addr
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " error: ", err)
			return
		}

		// accept the client
		fmt.Println("start Zinx server ", s.Name, "success, now listening...")

		// TODO 自动生成id
		var cid uint32
		cid = 0

		// 启动server网络连接业务，处理客户端连接请求
		for {
			// 阻塞等待客户端连接，处理客户端连接业务（读写）
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}

			// TODO Server.Start() 设置服务器的最大连接控制，如果超过最大连接，那么则关闭新的连接

			// TODO Server.Start() 处理该新连接请求的业务方法，此时应有 handler 和 conn 绑定
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("recv buf error: ", err)
			//			continue
			//		}
			//		// 回显
			//		if _, err := conn.Write(buf[:cnt]); err != nil {
			//			fmt.Println("write back buf error: ", err)
			//			continue
			//		}
			//	}
			//}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name ", s.Name)

	// TODO Server.Stop() 将其他需要清理的连接信息或者其他信息，也要一并停止或者清理
}

func (s *Server) Serve() {
	s.Start()

	// TODO

	for {
		time.Sleep(10 * time.Second)
	}
}

// 注册路由业务方法
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router success!")
}
