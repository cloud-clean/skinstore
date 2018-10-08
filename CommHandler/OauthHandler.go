package CommHandler

import (
	"net/http"
	"skinstore/common/logger"
)

var log = logger.NewLog()

func LotLoginHandler(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	redirectUrl := "https://open.bot.tmall.com/oauth/callback"
	acc := r.Form.Get("account")
	pwd := r.Form.Get("password")
	log.Infof("account:%s  password:%s",acc,pwd)
	referUrl := r.Header.Get("referer")
	log.Infof("referUrl:%s",referUrl)
	http.Redirect(w,r,redirectUrl,http.StatusMovedPermanently)
}
