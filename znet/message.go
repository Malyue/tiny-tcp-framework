package znet

type Message struct {
	Id   uint32
	Len  uint32
	Data []byte
}

func NewMessagePackage(id uint32, data []byte) *Message {
	return &Message{
		Id:   id,
		Len:  uint32(len(data)),
		Data: data,
	}
}

func (msg *Message) GetDataLen() uint32 {
	return msg.Len
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
	msg.setDataLen(uint32(len(msg.Data)))
}

func (msg *Message) setDataLen(len uint32) {
	msg.Len = len
}
