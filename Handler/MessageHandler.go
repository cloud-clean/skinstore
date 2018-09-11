package Handler

import (
	"net/http"
	"skinstore/WxService"
	"skinstore/common"
	"skinstore/web/router"
	"skinstore/Entity/wx/message"
	"strings"
	"skinstore/mqttClient"
)

func MsgGetHandler(params *router.Params,rw http.ResponseWriter) *common.WebResult {
	signature := params.Get("signature")
	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	echostr := params.Get("echostr")
	if WxService.CheckSignature(signature, nonce, timestamp) {
		rw.Write([]byte(echostr))
	} else {
		rw.Write([]byte{})
	}
	return nil
}

func MsgPostHandler(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	msg := params.GetWxMsg()
	log.Infof("type:%s  content:%s",msg.MsgType,msg.Content)
	go parseMsg(msg)
	rw.Write([]byte(""))
	return nil
}

func parseMsg(msg *message.WxMsg){
	if msg != nil{
		switch msg.MsgType{
		case message.MSG_TEXT:
			content := msg.Content
			if strings.Contains(content,":"){
				index := strings.Index(content,":")
				if index > 0{
					mqttClient.Mc.Publish(content)
				}
			}else{
				log.Info("get wx msg:%s  from %s",content,msg.FromUserName)
			}
		default:
			log.Info("do not parse this msg:%s",msg.MsgType)
		}
	}
}