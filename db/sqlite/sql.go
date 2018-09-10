package sqlite

func GetCreateSql(name string)string{
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