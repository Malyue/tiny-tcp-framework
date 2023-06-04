package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx/ziface"
	"zinx/zinx/utils"
)

// 封包拆包实例
type DataPack struct{}

// 初始化
func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}

// 封包方法（压缩数据）
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	databuffer := bytes.NewBuffer([]byte{})

	if err := binary.Write(databuffer, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(databuffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	if err := binary.Write(databuffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	fmt.Println(databuffer.Bytes())
	return databuffer.Bytes(), nil

}

// 解析head出来即可
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	dataBuffer := bytes.NewReader(binaryData)
	msg := &Message{}

	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断长度是否超出允许最大包长度
	if utils.GlobalCfg.MaxPacktSize > 0 && msg.Len > utils.GlobalCfg.MaxPacktSize {
		return nil, errors.New("Too large asg data recieved")
	}

	return msg, nil
}
