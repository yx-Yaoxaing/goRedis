package tcp

import (
	"context"
	"fmt"
	"goRedis/database"
	redis "goRedis/interface/database"
	"net"
	"sync"
)

type TcpConnectionHandler struct {
	activeConn sync.Map
	db         redis.DB
	closing    uint32
}

// 处理连接
func (tcpConnectionHandler *TcpConnectionHandler) Handle(ctx context.Context, conn net.Conn) {

	fmt.Println(conn.Read(make([]byte, 1024)))
}
func (tcpConnectionHandler *TcpConnectionHandler) Close() error {
	return nil
}

func MakeHandler() *TcpConnectionHandler {
	var db redis.DB
	// 初始化db实列  0-15个数据库
	db = database.NewStandaloneServer()
	return &TcpConnectionHandler{
		db: db,
	}
}
