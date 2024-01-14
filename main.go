package main

import (
	"fmt"
	"goRedis/config"
	"goRedis/tcp"
	"goRedis/util"
)

func main() {
	util.Print()
	// 默认配置
	config.Properties = config.Properties
	// 创建tcp连接
	tcp.ListenAndServe(&tcp.Config{
		Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
	}, tcp.MakeHandler())
}
