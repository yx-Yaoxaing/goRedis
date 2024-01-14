package database

import "goRedis/datastruct/dict"

const (
	dataDictSize = 1 << 16
)

// 数据处理
type DB struct {
	// 代表是第几个数据库 redis默认是0-15
	index int
	// 数据存储目录 完成k-v的映射 O(1)的查询新增
	data *dict.ConcurrentDict
}

func makeDB() *DB {
	db := &DB{
		data: dict.MakeConcurrent(dataDictSize),
	}
	return db
}
