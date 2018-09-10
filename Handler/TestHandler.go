package Handler

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"skinstore/WxService/WxConfig"
	"github.com/bitly/go-simplejson"
	"skinstore/Entity/wx"
	"skinstore/utils"
	"bytes"
)

func Test(rw http.ResponseWriter,req *http.Request){
	req.ParseForm()
	if req.Method == "GET"{
		db := utils.NewLevelDb()
		data := db.GetBytes(fmt.Sprintf("%s%s",WxConfig.STORE_PREFIX_USERINFO,"cloud"))
		reader := bytes.NewReader(data)
		user := &wx.WxUserInfo{}
		user.Deserialize(reader)
		rw.Write(data)

	}else if req.Method == "POST"{
		body,_:= ioutil.ReadAll(req.Body)
		json,err := simplejson.NewJson(body)
		if err != nil{
			log.Error(err)
			rw.WriteHeader(404)
		}else{
			user := &wx.WxUserInfo{
				Openid:json.Get("openid").MustString(),
				Nickname:json.Get("nickname").MustString(),
				City:json.Get("city").MustString(),
				Sex:json.Get("sex").MustString(),
				Country:json.Get("country").MustString(),
				Headimgurl:json.Get("icon").MustString(),
				Province:json.Get("provience").MustString(),
			}
			db := utils.NewLevelDb()
			data := new(bytes.Buffer)
			user.Serialize(data)
			db.SetBytes(fmt.Sprintf("%s%s",WxConfig.STORE_PREFIX_USERINFO,"cloud"),data.Bytes())
		}

	}
}
