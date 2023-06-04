package ziface

// 原先的请求只有一个[]byte数组作为消息，但是太简单了，需要再做一层封装

type IMessage interface {
	// 消息长度
	GetDataLen() uint32
	// 消息ID
	GetMsgId() uint32
	// 消息内容
	GetData() []byte
	// 设计消息ID/内容/长度
	SetMsgId(uint32)
	SetData([]byte)
	//SetDataLen(uint32)
}
