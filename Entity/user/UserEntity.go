package user

import (
	"time"
	"skinstore/utils/SqliteUtil"
	"skinstore/common"
	"errors"
)

type UserEntity struct {
	UserId int				`json:"userId"`
	OpenId string			`json:"openId"`
	Nickname string			`json:"nickName"`
	Mobile string			`json:"mobile"`
	ImgUrl string			`json:"imgUrl"`
	CreateTm time.Time		`json:"createTm"`
}

func (user *UserEntity) Save() error{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("insert into user(open_id,nick_name,mobile,img_url) values(?,?,?,?);")
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Exec(user.OpenId,user.Nickname,user.Mobile,user.ImgUrl)
	common.CheckErr(err)
	if rows,err := res.RowsAffected();rows < 1 || err != nil{
		return errors.New("save user to sql fail")
	}
	return nil
}

func (user *UserEntity)GetUserInfo(){
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare(`select user_id as UserId,nick_name as NickName, img_url as ImgUrl 
					from user where open_id = ? limit 1`)
	defer stmt.Close()
	common.CheckErr(err)
	res,err := stmt.Query(user.OpenId)
	common.CheckErr(err)
	for res.Next(){
		err = res.Scan(user.UserId,user.Nickname,user.ImgUrl)
		common.CheckErr(err)
	}
}

func IsLoginUser(userId,openId string)bool{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("select count(0) from user where user_id = ? and open_id = ?")
	common.CheckErr(err)
	defer stmt.Close()
	rows := stmt.QueryRow(userId,openId)
	var count int
	rows.Scan(&count)
	if count == 0 {
		return true
	}else{
		return false
	}

}


