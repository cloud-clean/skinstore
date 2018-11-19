package lot

import (
	"skinstore/common"
	"skinstore/utils/SqliteUtil"
	"time"
)

type AuthToken struct {
	Account string `json:'account'`
	AccessToken string `json:'accessToken'`
	FlashToken string `json:'flashToken'`
	Expire time.Time `json:'expire'`
}

type AuthCode struct {
	Account string `json:'account'`
	Code string `json:'code'`
	Expire time.Time `json:'expire`
}

func (this *AuthCode) update() error{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("replace into auth_code(`account`,`code`,`expire`) values(?,?,?)")
	common.CheckErr(err)
	_,err = stmt.Exec(this.Account,this.Code,this.Expire)
	return err
}

func ConfirmCode(code string) bool{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select count(1) from auth_code where code = ? and expire > now()")
	common.CheckErr(err)
	var count = 0
	res := stmt.QueryRow(code)
	err = res.Scan(&count)
	if err ==nil{
		if count>0{
			return true
		}
	}
	return false
}
