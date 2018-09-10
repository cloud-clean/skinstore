package wx

import (
	"testing"
	"bytes"
	"fmt"
)

func TestWxUserInfo_Deserialize(t *testing.T) {
	u := &WxUserInfo{
		Openid:"dddd",
		City:"city",
		Country:"country",
		Nickname:"nickname",
		Sex:"s",
		Province:"province",
		Headimgurl:"url",
		Unionid:"unioned",
	}
	by := new(bytes.Buffer)
	u.Serialize(by)

	fmt.Println(by)
	info := new(WxUserInfo)
	info.Deserialize(by)
	fmt.Println(info.Unionid)

}
