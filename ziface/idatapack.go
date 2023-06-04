package ziface

/*
	解决粘包问题
	面向TCP连接中的数据流，为传输数据添加头部信息，用于处理TCP粘包问题
*/

type IDataPack interface {
	// 获取包长度方法
	GetHeadLen() uint32
	// 封包方法
	Pack(msg IMessage) ([]byte, error)
	// 拆包方法
	Unpack([]byte) (IMessage, error)
}
