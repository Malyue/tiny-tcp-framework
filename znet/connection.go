package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 连接的ID,即SessionID
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool

	// 当前连接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc

	// 该连接的处理方法Router
	Router ziface.IRouter

	// 告知当前连接已经退出/停止的channel
	ExitBuffChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 创建拆包解包的对象
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 拆包，得到msgid和len
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			c.ExitBuffChan <- true
			continue
		}

		// 根据len读取data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				c.ExitBuffChan <- true
				continue
			}
		}

		msg.SetData(data)

		// 得到当前客户端请求的request数据
		req := Request{
			c,
			msg,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		//// 读取最大的数据到buf中
		//buf := make([]byte, 512)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf error: ", err)
		//	c.ExitBuffChan <- true
		//	continue
		//}
		//// 得到当前客户端请求的Request数据
		//req := Request{
		//	conn: c,
		//	msg:,
		//}
		//// 从路由Routers中找到注册绑定的Conn的对应Handle
		//go func(request ziface.IRequest) {
		//	//注册执行的路由方法
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)
		// 调用当前链接业务
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnID ", c.ConnID, " handle is error: ", err)
		//	c.ExitBuffChan <- true
		//	continue
		//}
	}
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("conn closed when send msg")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMessagePackage(msgId, data))
	if err != nil {
		fmt.Println("pack err msg id = ", msgId)
		return errors.New("pack err msg : " + err.Error())
	}

	// 写回客户端
	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("Write msg id ", msgId, " error")
		c.ExitBuffChan <- true
		return err
	}

	return nil
}

func (c *Connection) Start() {
	// 开启处理该链接读取到客户端数据之后的请求业务
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			//得到退出消息，不再阻塞
			return
		}
	}
}

// 停止链接，结束当前连接状态M
func (c *Connection) Stop() {
	// 如果连接已关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// TODO Connection Stop()

	// 关闭socket
	c.Conn.Close()

	// 通知从缓冲队列读取数据的业务，该连接已关闭
	c.ExitBuffChan <- true

	// 关闭该连接全部管道
	close(c.ExitBuffChan)
}

// 从当前连接获取原始的 socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
