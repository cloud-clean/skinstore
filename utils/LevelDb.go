package utils

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"fmt"
	"sync"
)

type  LevelDb struct{
	db *leveldb.DB
}
var store *LevelDb
var once sync.Once



func NewLevelDb()*LevelDb {
	once.Do(func(){
		db,err := leveldb.OpenFile("./db/levelDb",nil)
		if err != nil {
			log.Fatal("open db fail")
		}
		store = &LevelDb{db}
	})
	return store
}

func (db *LevelDb)Get(key string) string {
	res,err := db.db.Get([]byte(key),nil)
	if err != nil {
		log.Println(fmt.Sprint("leveldb get key:%s fail",key))
		return ""
	}
	return string(res)
}

func (db *LevelDb)GetBytes(key string) []byte{
	res,err := db.db.Get([]byte(key),nil)
	if err != nil {
		log.Println(fmt.Sprint("leveldb get key:%s fail",key))
		return nil
	}
	return res
}

func (db *LevelDb)SetString(key,value string) {
	db.db.Put([]byte(key),[]byte(value),nil)
}
func (db *LevelDb)SetBytes(key string,value []byte) {
	db.db.Put([]byte(key),value,nil)
}

func (db *LevelDb) Del(key string){
	db.db.Delete([]byte(key),nil)
}

func (db *LevelDb) Close(){
	db.db.Close()
}




