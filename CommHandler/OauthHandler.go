package CommHandler

import (
	"encoding/json"
	"github.com/silenceper/wechat/util"
	"net/http"
	"net/url"
	"skinstore/common"
	"skinstore/common/logger"
	"skinstore/mqttClient"
	"strings"
	"io/ioutil"
	"skinstore/Entity/lot"
	"time"
	"regexp"
)

var log = logger.NewLog()

func LotLoginHandler(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	acc := r.Form.Get("account")
	pwd := r.Form.Get("password")
	reg := regexp.MustCompile(`https:.*skillId=(\d{5,})`)
	log.Infof("account:%s  pwd:%s",acc,pwd);
	if !lot.Login(acc,pwd){
		w.WriteHeader(http.StatusForbidden)
		return;
	}
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
	var authCode = &lot.AuthCode{Account:acc,Code:util.RandomStr(16),Expire:time.Now().Add(10*time.Minute)}
	err := authCode.Update()
	if err == nil{
		log.Infof("account:%s  password:%s  skillId:%s   clientId:%s",acc,pwd,skillId,clientId)
		redirectUrl = redirectUrl+"?code="+authCode.Code+"&state="+state+"&token="+token+"&client_id="+clientId+"&skillId="+skillId
		log.Infof("redirectUrl:%s",redirectUrl)
		http.Redirect(w,r,redirectUrl,http.StatusMovedPermanently)
	}else{
		log.Error(err.Error())
		w.WriteHeader(http.StatusNotImplemented)
	}


}

func LotTokenAccess(r *http.Request,w http.ResponseWriter){
	r.ParseForm()
	clientId := r.PostForm.Get("client_id")
	clientSecret:= r.PostForm.Get("client_secret")
	code := r.PostForm.Get("code")
	log.Infof("code:%s",code)
	grantType := r.PostForm.Get("grant_type")
	var resp = make(map[string]interface{})
	if account,err := lot.ConfirmCode(code);err == nil{
		log.Infof("clientId:%s  clientSecret:%s  code:%s  grantType:%s",clientId,clientSecret,code,grantType)
		var acesstoken = &lot.AuthToken{
			Account:account,
			AccessToken:util.RandomStr(18),
			FlashToken:util.RandomStr(16),
			Expire:time.Now().Add(17600*time.Second),
		}
		err := acesstoken.Update()
		if err == nil {
			resp["access_token"] = acesstoken.AccessToken
			resp["refresh_token"] = acesstoken.FlashToken
			resp["expires_in"] = 17600000
			w.Header().Set("Content-Type","application/json")
			b,err := json.Marshal(resp)
			common.CheckErr(err)
			w.Write(b)
		}else {
			resp["error"] = "operate fail"
			w.Header().Set("Content-Type", "application/json")
			b, _ := json.Marshal(resp)
			w.Write(b)
		}
	}else{
		resp["error"] = "state error"
		w.Header().Set("Content-Type","application/json")
		b,_ := json.Marshal(resp)
		w.Write(b)
	}



}

func LotCallback(r *http.Request,w http.ResponseWriter){
	if len(r.Header) > 0{
		//log.Info("headers:")
		//for k,v := range r.Header{
		//	log.Infof("%s:%s",k,v)
		//}
		res,_ := ioutil.ReadAll(r.Body)
		var callParam lot.AliCallback
		log.Info(string(res))
		json.Unmarshal(res,&callParam)
		log.Info(callParam.Header.Namespace)
		accessToken := callParam.Payload.AccessToken
		log.Info(accessToken)
		log.Info(callParam.Header.Namespace)
		switch(callParam.Header.Namespace){
		case "AliGenie.Iot.Device.Discovery":
			var resp lot.AliCallback
			resp.Header.Namespace = "AliGenie.Iot.Device.Discovery"
			resp.Header.Name = "DiscoveryDevicesResponse"
			resp.Header.MessageId = callParam.Header.MessageId
			resp.Header.PayLoadVersion = 1
			devices := []lot.LotDevice{lot.LotDevice{DeviceId:"cloud_lot_1",DeviceName:"lightLot",DeviceType:"light",Model:"lot",Brand:"cloud",Icon:"https://www.home-assistant.io/demo/favicon-192x192.png"}}
			var proMap = make(map[string]string)
			proMap["status"] = "off"
			devices[0].Properties = []map[string]string{proMap}
			resp.Payload.Devices = devices
			resp.Payload.Actions = []string{"TurnOn","TurnOff"}
			b,err := json.Marshal(resp)
			log.Info(string(b))
			if err != nil{

			}
			w.Header().Set("Content-Type","application/json")
			w.Write(b)
			log.Info("finish")
			case "AliGenie.Iot.Device.Control":
				var resp lot.AliCallback
				resp.Header.Namespace = "AliGenie.Iot.Device.Control"
				resp.Header.MessageId = callParam.Header.MessageId
				resp.Header.PayLoadVersion = 1
				resp.Payload.DeviceId = callParam.Payload.DeviceId
				if "TurnOn" == callParam.Header.Name{
					log.Info("打开灯")
					var msg = lot.LampStatusEntity{Pos:"lot3",Status:"1"}
					mqttClient.SendMsg(msg)
					resp.Header.Name = "TurnOnResponse"

				}else{
					log.Info("关闭灯")
					var msg = lot.LampStatusEntity{Pos:"lot3",Status:"0"}
					mqttClient.SendMsg(msg)
					resp.Header.Name = "TurnOffResponse"
				}
				b,_ := json.Marshal(resp)
				w.Header().Set("Content-Type","application/json")
				w.Write(b)
		}

	}

}
