package main

import (
	"fmt"
	"github.com/cznic/kv"
	"log"
	"os"
)

// 打开或者创建KV数据库
// 当path指向的数据库存在时打开该数据库，否则尝试在该路径处创建新数据库
func OpenOrCreateKv(path string, options *kv.Options) (*kv.DB, error) {
	db, errOpen := kv.Open(path, options)
	if errOpen != nil {
		var errCreate error
		db, errCreate = kv.Create(path, options)
		if errCreate != nil {
			return db, errCreate
		}
	}

	return db, nil
}

func main() {
	dbPath := "test.kv"
	db, err := OpenOrCreateKv(dbPath, &kv.Options{})
	if db == nil || err != nil {
		log.Fatal("无法打开数据库", dbPath, ": ", err)
	}
	defer db.Close()

	db.Set([]byte("key1"), []byte("value1"))
	db.Set([]byte("key1"), []byte("value2"))

	buffer, _ := db.Get(nil, []byte("key1"))
	fmt.Printf("%s\n", buffer)

	walFile := db.WALName()
	db.Close()
	os.Remove(walFile)
	os.Remove("test.kv")
}
