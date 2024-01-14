package tcp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"
)

// 实现通信tcp

// 定义tcp所需数据结构
type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}

func ListenAndServe(cfg *Config, handler Handler) error {
	// 监听端口 tcp
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("bind: %s, start listening...", cfg.Address))
	ctx := context.Background()
	// 监听客户端请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			return errors.New("listener.Accept() 失败")
		}
		// 开启协程 处理这个连接
		go handler.Handle(ctx, conn)
	}
	return nil
}
