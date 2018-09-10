package WxService

import (
	"skinstore/WxService/WxConfig"
	"sort"
	"strings"
	"crypto/sha1"
	"encoding/hex"
	"skinstore/utils/httpUtil"
	"log"
	"skinstore/Entity/wx"
)

func CheckSignature(signature,nonce,timestamp string) bool{
	sorts := []string{WxConfig.Token,timestamp,nonce}
	sort.Strings(sorts)
	sortStr := strings.Join(sorts,"")
	sha := sha1.New()
	sha.Write([]byte(sortStr))
	sign := sha.Sum(nil)
	if signature == hex.EncodeToString(sign){
		return true
	}else{
		return false
	}
}

func GetOpenId(code string) (string,string,string){
	params := make(map[string]interface{})
	params["code"] = code
	params["appid"] = WxConfig.AppID
	params["secret"] =WxConfig.AppSecret
	params["grant_type"] = "authorization_code"
	res := httpUtil.JsonGet(WxConfig.WxOpenIdUrl,params)
	accToken := res.Get("access_token").MustString()
	openId := res.Get("openid").MustString()
	lang := res.Get("lang").MustString()
	if openId == ""{
		errCode := res.Get("errcode").MustInt()
		errmsg := res.Get("errmsg").MustString()
		log.Printf("get weixin openid fail code:%d msg:%s",errCode,errmsg)
		return "","",""
	}
	return openId,accToken,lang
 }

 func GetUserInfo(openId,accessToken,lang string) *wx.WxUserInfo{
	params := make(map[string]interface{})
	params["access_token"] = accessToken
	params["openid"] = openId
	params["lang"] = lang
	res := httpUtil.JsonGet(WxConfig.WxUserInfoUrl,params)
	if res == nil {
		log.Printf("get wx user info fail")
		return nil
	}
	errcode := res.Get("errcode").MustInt()
	if errcode >0 {
		log.Printf("get wx user info err:%s",res.Get("errmsg").MustString())
		return nil
	}
	userinfo := &wx.WxUserInfo{
		Openid:openId,
		Nickname:res.Get("nickname").MustString(),
		Sex:res.Get("sex").MustString(),
		Province:res.Get("province").MustString(),
		City:res.Get("city").MustString(),
		Country:res.Get("country").MustString(),
		Headimgurl:res.Get("headimgurl").MustString(),
		Unionid:res.Get("unionid").MustString(),
		Privilege:res.Get("privilege").MustStringArray(),
	}
	return userinfo
 }

