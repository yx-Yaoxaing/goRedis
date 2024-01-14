package tcp

import (
	"context"
	"net"
)

type HandleFunc func(ctx context.Context, conn net.Conn)

// 处理器定义接口
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
