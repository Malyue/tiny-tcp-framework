package utils

import (
	"encoding/json"
	"log"
	"os"
	"zinx/ziface"
)

const configPath = "conf/zinx.json"

type GlobalConfig struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	Version   string

	// 数据包的最大值
	MaxPacktSize uint32
	// 允许最大的连接个数
	MaxConn int
}

var GlobalCfg *GlobalConfig

func init() {
	// 设置默认值
	GlobalCfg = &GlobalConfig{
		Name:         "ZinxServerApp",
		Version:      "V0.1",
		TcpPort:      7777,
		Host:         "0.0.0.0",
		MaxConn:      12000,
		MaxPacktSize: 4096,
	}
	// 从配置文件中加载
	GlobalCfg.Reload()
}

func (g *GlobalConfig) Reload() {
	exists, err := PathExists(configPath)
	if err != nil {
		panic(err)
	}
	if !exists {
		log.Printf("The config file %s is not exist,use the default config", configPath)
		return
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalCfg)
	if err != nil {
		panic(err)
	}
}

func GetGlobalConfig() *GlobalConfig {
	if GlobalCfg != nil {
		return GlobalCfg
	}
	GlobalCfg.Reload()
	return GlobalCfg
}
