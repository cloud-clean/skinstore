package SqliteUtil

import (
	"sync"
	"database/sql"
	"skinstore/common"
	"os"
	"path"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	Db *sql.DB
}


var instance *SqliteDb
var once sync.Once


func NewSqlDb() *SqliteDb{
	once.Do(func(){
		//check file is exists
		filename := "./db/sqlite/skinstore.db"
		_,err := os.Stat(filename)
		if err != nil && os.IsNotExist(err){
			err := os.MkdirAll(path.Dir(filename),os.ModePerm)
			common.CheckErr(err)
			_,err=os.Create(filename)
			common.CheckErr(err)
		}
		db,err := sql.Open("sqlite3",filename)
		common.CheckErr(err)
		//check all table is exists,if not create
		sqlTables := map[string]string{
			"user":"0",
			"reservation":"0",
			"project":"0",
		}
		tables,err := db.Query("select name from sqlite_master")
		common.CheckErr(err)
		defer tables.Close()
		for tables.Next(){
			var table string
			tables.Scan(&table)
			if _,ok := sqlTables[table];ok{
				sqlTables[table] = "1"
			}
		}
		for k,v := range sqlTables {
			if v == "0"{
				sqlStr := getCreateSql(k)
				_,err := db.Exec(sqlStr)
				common.CheckErr(err)

			}
		}
		instance = &SqliteDb{Db:db}
	})
	return instance
}

func (db *SqliteDb) Close(){
	db.Db.Close()
}


func ResultMap(res sql.Result,t interface{}){
	//cols,_ := res.
}

func getCreateSql(name string)string{
	switch name {
	case "user":
		return "CREATE TABLE user (user_id INTEGER PRIMARY KEY ON CONFLICT ROLLBACK AUTOINCREMENT, open_id VARCHAR (32) UNIQUE NOT NULL, nick_name VARCHAR (50), mobile VARCHAR (12), create_tm DATETIME DEFAULT ((datetime('now', 'localtime'))) COLLATE RTRIM, img_url VARCHAR (200));"
	case "reservation":
		return "CREATE TABLE reservation (id INTEGER PRIMARY KEY AUTOINCREMENT,user_id INTEGER, name VARCHAR (50) NOT NULL,re_time DATETIME NOT NULL,status CHAR (1) DEFAULT (0),mobile VARCHAR (12) NOT NULL,create_tm DATETIME DEFAULT (datetime('now', 'localtime')));CREATE INDEX INDEX_TIME ON resevation (re_time DESC);"
	case "project":
		return "CREATE TABLE project (id INTEGER PRIMARY KEY AUTOINCREMENT,name VARCHAR (50) NOT NULL,description VARCHAR (500),original_price INTEGER,type VARCHAR (5) NOT NULL,img_url VARCHAR (200) NOT NULL,cur_price INTEGER NOT NULL,status CHAR (1) DEFAULT (0));"
	default:
		return ""
	}
}