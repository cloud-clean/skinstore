package Handler

import (
	"net/http"
	"skinstore/WxService"
	user "skinstore/Entity/user"
	"skinstore/common"
	"skinstore/web/router"
)

func WxOauthHandler(param *router.Params,rw http.ResponseWriter)*common.WebResult{
	code := param.Get("code")
	state := param.Get("state")
	openid,acc,lang := WxService.GetOpenId(code)
	if openid != ""{
		wxUser := WxService.GetUserInfo(openid,acc,lang)
		if wxUser != nil {
			userEntity := &user.UserEntity{
				OpenId:wxUser.Openid,
				Nickname:wxUser.Nickname,
				ImgUrl:wxUser.Headimgurl,
				Mobile:"",
			}
			err := userEntity.Save()
			common.CheckErr(err)
		}
		cookie := http.Cookie{Name:"openid",Value:openid,Path:"/",MaxAge:86400*7}
		http.SetCookie(rw,&cookie)
		http.Redirect(rw,&http.Request{Method:"GET"},state,http.StatusFound)
	}
	return nil
}
