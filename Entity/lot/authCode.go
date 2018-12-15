package lot

import (
	"errors"
	"github.com/silenceper/wechat/util"
	"skinstore/common"
	"skinstore/utils/SqliteUtil"
	"time"
	"fmt"
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

func (this *AuthCode) Update() error{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("replace into auth_code(`account`,`code`,`expire`) values(?,?,?)")
	common.CheckErr(err)
	defer stmt.Close()
	_,err = stmt.Exec(this.Account,this.Code,this.Expire)
	return err
}

func ListCode() []AuthCode{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select account,code,expire from auth_code")
	defer stmt.Close()
	common.CheckErr(err)
	res,err := stmt.Query()
	common.CheckErr(err)
	var list []AuthCode
	for res.Next(){
		var authCode AuthCode
		err = res.Scan(&authCode.Account,&authCode.Code,&authCode.Expire)
		list = append(list, authCode)
	}
	return list;
}

func ListAccessTokens() []AuthToken{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select account,access_token,flash_token,expire from auth_token")
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Query()
	common.CheckErr(err)
	var list []AuthToken
	for res.Next(){
		var access AuthToken
		err = res.Scan(&access.Account,&access.AccessToken,&access.FlashToken,&access.Expire)
		list = append(list, access)
	}
	return list;
}

func (this *AuthToken) Update() error{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("replace into access_token(`account`,`access_token`,`flash_token`,`expire`) values(?,?,?,?)")
	common.CheckErr(err)
	defer stmt.Close()
	_,err = stmt.Exec(this.Account,this.AccessToken,this.FlashToken,this.Expire)
	return err
}

func ConfirmCode(code string) (string,error){
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select account from auth_code where code = ? and expire > datetime('now','localtime')")
	defer stmt.Close()
	common.CheckErr(err)

	var account = ""
	res := stmt.QueryRow(code)
	err = res.Scan(&account)
	fmt.Println("account:"+account)
	if err ==nil{
		if account != ""{
			return account,nil
		}
	}
	return "",errors.New("code is not exists")
}

func fleshCode(code string) (*AuthToken,error){
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select account,assess_token as accessToken from auth_token where flash_token = ?")
	defer stmt.Close()
	common.CheckErr(err)
	var token = &AuthToken{}
	res := stmt.QueryRow(code)
	err = res.Scan(token.Account,token.AccessToken)
	if err ==nil{
		if token.Account != ""{
			token.AccessToken = util.RandomStr(18)
			token.Expire = time.Now().Add(17600000*time.Microsecond)
			token.FlashToken = util.RandomStr(18)
			token.Update()
			return token,nil
		}
	}
	return nil,errors.New("flas token is error")
}

