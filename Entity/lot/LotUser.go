package lot

import (
	"skinstore/common"
	"skinstore/common/logger"
	"skinstore/utils/SqliteUtil"
)
var log = logger.NewLog()

type LotUser struct {
	Account string `json:'account'`
	Passworld string `json:'password'`
	Group string `json:'group'`
}


func Login(account,password string)bool{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select count(1) from lot_user where account = ? and password = ?")
	common.CheckErr(err)
	defer stmt.Close()
	res := stmt.QueryRow(account,password)
	count:=0
	err = res.Scan(&count)
	if err != nil{
		log.Error(err.Error())
		return false;
	}
	if count == 1{
		return true;
	}
	return  false;
}

func (this *LotUser)Save() error {
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("insert into lot_user(`account`,`password`,`group`) values(?,?,?)")
	common.CheckErr(err)
	defer stmt.Close()
	_,err = stmt.Exec(this.Account,this.Passworld,this.Group)
	return err;
}
