package CommHandler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"skinstore/common"
	"skinstore/common/logger"
	"strings"
)

var log = logger.NewLog()

func LotLoginHandler(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	acc := r.Form.Get("account")
	pwd := r.Form.Get("password")
	reg := regexp.MustCompile(`https:.*skillId=(\d{5,})`)

	referUrl := r.Header.Get("referer")
	decoderUrl,_ := url.QueryUnescape(referUrl)
	referUri,_ := url.ParseRequestURI(decoderUrl)
	refParams := referUri.Query()
	redirectUrl := refParams.Get("redirect_uri")
	log.Infof("redicect:%s",redirectUrl)
	skillId := reg.FindStringSubmatch(redirectUrl)[1]
	redirectUrl = strings.Split(redirectUrl,"?")[0]
	clientId := refParams.Get("client_id")
	token:=refParams.Get("token")
	state := refParams.Get("state")
	//resType := refParams.Get("response_type")
	log.Infof("account:%s  password:%s  skillId:%s   clientId:%s",acc,pwd,skillId,clientId)
	redirectUrl = redirectUrl+"?code=asdfaeafaea&state="+state+"&token="+token+"&client_id="+clientId+"&skillId="+skillId
	log.Infof("redirectUrl:%s",redirectUrl)
	http.Redirect(w,r,redirectUrl,http.StatusMovedPermanently)
}

func LotTokenAccess(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	clientId := r.PostForm.Get("client_id")
	clientSecret:= r.PostForm.Get("client_secret")
	code := r.PostForm.Get("code")
	grantType := r.PostForm.Get("grant_type");
	log.Infof("clientId:%s  clientSecret:%s  code:%s  grantType:%s",clientId,clientSecret,code,grantType)
	var resp = make(map[string]interface{})
	resp["access_token"] = "xxxxdfasdfa"
	resp["refresh_token"] = "eaesfasefa"
	resp["expires_in"] = 17600000
	w.Header().Set("Content-Type","application/json")
	b,err := json.Marshal(resp)
	common.CheckErr(err)
	w.Write(b)

}
