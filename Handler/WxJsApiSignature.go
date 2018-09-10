package Handler

import (
	"net/http"
	"time"
	"fmt"
	"github.com/satori/go.uuid"
	"skinstore/WxService/WxConfig"
	"sort"
	"strings"
	"crypto/sha1"
	"encoding/hex"
	"skinstore/common"
	"skinstore/web/router"
)

func JsapiSignHandler(params *router.Params,rw http.ResponseWriter)*common.WebResult{
	var array []string
	array = append(array,fmt.Sprintf("%s=%s","url",params.Get("url")))
	timestamp := time.Now().Unix()
	array = append(array,fmt.Sprintf("%s=%d","timestamp",timestamp))
	nonce,err := uuid.NewV4()
	if err != nil {
		log.Error("uuid fail")
	}
	array = append(array,fmt.Sprintf("%s=%s","noncestr",nonce.String()))
	jsapiticket := WxConfig.GetJsTicket()
	array = append(array,fmt.Sprintf("%s=%s","jsapi_ticket",jsapiticket))
	sort.Strings(array)
	str:=strings.Join(array,"&")
	recode := sha1.Sum([]byte(str))
	signal := hex.EncodeToString(recode[:])
	res := make(map[string]interface{})
	res["signature"] = signal
	res["timestamp"] = timestamp
	res["nonceStr"] = nonce.String()
	res["appId"] = WxConfig.AppID
	return common.NewResult(1,res)
}
