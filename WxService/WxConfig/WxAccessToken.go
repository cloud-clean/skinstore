package WxConfig

import (
	"time"
	"skinstore/utils/httpUtil"
	"sync"
	"skinstore/utils/SqliteUtil"
	"skinstore/common"
	"skinstore/common/logger"
)

var log = logger.NewLog()
type WxAccessToken struct {
	accessToken string
	expiresIn time.Time
	lock sync.Mutex
}

type JsApiTicket struct {
	Ticket string
	expiresIn time.Time
	lock sync.Mutex
}

var ticket *JsApiTicket = &JsApiTicket{
	Ticket:"",
	expiresIn:time.Now(),
	lock:sync.Mutex{},
}

var token *WxAccessToken = &WxAccessToken{
	accessToken:"",
	expiresIn:time.Now(),
	lock:sync.Mutex{},
}



func GetWxAccessToken() string{
	token.lock.Lock()
	defer token.lock.Unlock()
	if token.isExpire(){
		freshAccessToken()
		return token.accessToken
	}else{
		return token.accessToken
	}
}

func (at *WxAccessToken) isExpire() bool{
	now := time.Now()
	if now.After(at.expiresIn){
		return true
	}else{
		return false
	}
}

func (at *WxAccessToken)updateAccessInDb(){
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("update wx set token = ? and expire = ? where key = 'ACCTOKEN")
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Exec(at.accessToken,at.expiresIn)
	common.CheckErr(err)
	int,err := res.RowsAffected()
	common.CheckErr(err)
	if int < 0{
		log.Error("update wx access token failed")
	}
}

func (t *JsApiTicket) isExpire() bool{
	if t == nil{
		return true
	}
	now := time.Now()
	if now.After(t.expiresIn){
		return true
	}else{
		return false
	}
}

func (t *JsApiTicket)updateJsTicketInDb(){
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("update wx set token = ? and expire = ? where key = 'JSTOKEN")
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Exec(t.Ticket,t.expiresIn)
	common.CheckErr(err)
	int,err := res.RowsAffected()
	common.CheckErr(err)
	if int < 0{
		log.Error("update wx jsticket token failed")
	}
}

func freshAccessToken() {
	params := make(map[string]interface{})
	params["grant_type"]="client_credential"
	params["appid"] = AppID
	params["secret"] = AppSecret
	res := httpUtil.JsonGet(WxAccessTokenUrl,params)
	accToken := res.Get("access_token").MustString()
	if accToken == "" {
		errCode := res.Get("errcode").MustInt()
		errMsg := res.Get("errmsg").MustString()
		log.Infof("get accessToken error code:%d msg: %s",errCode,errMsg)
		return
	}
	token = &WxAccessToken{
		accessToken:accToken,
		expiresIn:time.Now().Add(time.Second*7150),
	}
	token.updateAccessInDb()
}

func freshJsTicket(){
	params := make(map[string]interface{})
	params["access_token"] = GetWxAccessToken()
	params["type"] = "jsapi"
	res := httpUtil.JsonGet(WxJsTicketUrl,params)
	errcode := res.Get("errcode").MustInt()
	if errcode == 0{
		ticket = &JsApiTicket{
			Ticket:res.Get("ticket").MustString(),
			expiresIn:time.Now().Add(time.Second*7150),
		}
		ticket.updateJsTicketInDb()
	}else{
		log.Infof("get wx jsapi ticket fail msg:%s",res.Get("errmsg").MustString())
	}
}

func GetJsTicket() string{
	ticket.lock.Lock()
	defer ticket.lock.Unlock()
	if ticket.isExpire(){
		freshJsTicket()
		return ticket.Ticket
	}else{
		return ticket.Ticket
	}
}




