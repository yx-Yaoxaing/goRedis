package database

import (
	"goRedis/config"
	"goRedis/interface/redis"
	"sync/atomic"
)

type Server struct {
	dbSet []*atomic.Value // *DB

}

func (server *Server) Exec(client redis.Connection, cmdLine [][]byte) redis.Reply {
	return nil
}

func (server *Server) AfterClientClose(c redis.Connection) {

}
func (server *Server) Close() {

}

func NewStandaloneServer() *Server {
	server := &Server{}
	if config.Properties.Databases == 0 {
		// 就跟Redis一样 默认16个库
		config.Properties.Databases = 16
	}
	server.dbSet = make([]*atomic.Value, config.Properties.Databases)
	// 初始化db start
	for i := range server.dbSet {
		// 每一个数据库实列 就是对应一个dict
		singleDB := makeDB()
		singleDB.index = i
		holder := &atomic.Value{}
		holder.Store(singleDB)
		server.dbSet[i] = holder
	}
	return server
}
