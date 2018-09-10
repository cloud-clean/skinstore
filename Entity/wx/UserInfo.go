package wx

import (
	"io"
	"skinstore/common/serialize"
)

type WxUserInfo struct {
	Openid string		`json:"openid"`
	Nickname string		`json:"nickname"`
	Sex string			`json:"sex"`
	Province string		`json:"province"`
	City string			`json:"city"`
	Country string		`json:"country"`
	Headimgurl string	`json:"headimgurl"`
	Unionid string		`json:"unionid"`
	Privilege []string	`json:"privilege"`
}

func (u *WxUserInfo) Serialize(w io.Writer) error{
	err:=serialize.WriteString(w,u.Openid)
	err = serialize.WriteString(w,u.Nickname)
	err = serialize.WriteString(w,u.Sex)
	err = serialize.WriteString(w,u.Province)
	err = serialize.WriteString(w,u.City)
	err = serialize.WriteString(w,u.Country)
	err = serialize.WriteString(w,u.Headimgurl)
	err = serialize.WriteString(w,u.Unionid)
	return err
}

func (u *WxUserInfo) Deserialize(r io.Reader)error{
	openid,err := serialize.ReadString(r)
	nickname,err := serialize.ReadString(r)
	sex,err := serialize.ReadString(r)
	Province,err := serialize.ReadString(r)
	City,err := serialize.ReadString(r)
	Country,err := serialize.ReadString(r)
	Headimgurl,err := serialize.ReadString(r)
	Unionid,err := serialize.ReadString(r)
	u.Unionid = Unionid
	u.Openid = openid
	u.Nickname = nickname
	u.Sex = sex
	u.Province = Province
	u.City = City
	u.Country = Country
	u.Headimgurl = Headimgurl
	if err != nil {
		return err
	}else{
		return nil
	}

}
