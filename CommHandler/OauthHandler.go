package CommHandler

import (
	"net/http"
	"net/url"
	"skinstore/common/logger"
)

var log = logger.NewLog()

func LotLoginHandler(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	acc := r.Form.Get("account")
	pwd := r.Form.Get("password")

	referUrl := r.Header.Get("referer")
	decoderUrl,_ := url.QueryUnescape(referUrl)
	referUri,_ := url.ParseRequestURI(decoderUrl)
	refParams := referUri.Query()
	redirectUrl := refParams.Get("redirect_uri")
	skillId := refParams.Get("skillId")
	clientId := refParams.Get("client_id")
	state := refParams.Get("state")
	//resType := refParams.Get("response_type")
	log.Infof("account:%s  password:%s  skillId:%s   clientId:%s",acc,pwd,skillId,clientId)
	redirectUrl = redirectUrl+"?toke=asdfaeafaea&state="+state
	log.Infof("redirectUrl:%s",redirectUrl)
	http.Redirect(w,r,redirectUrl,http.StatusMovedPermanently)
}

func LotTokenAccess(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	for k,v := range r.PostForm{
		log.Infof("key:%s  value:%s",k,v)
	}
}
