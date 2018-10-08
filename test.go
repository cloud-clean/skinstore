package main

import (
	"crypto"
	"encoding/hex"
	"fmt"
)

func main(){
	//str := "cmd:haha"
	//fmt.Println(str[:strings.Index(str,":")])
	//fmt.Println(str[strings.Index(str,":")+1:])
	//tcp.Start()
	key:="appkey";
	pass:="123456";
	encode:="2126622d53e62f60f8e4b23358a218a1";
	encoder := crypto.MD5.New()
	encoder.Write([]byte(key+pass))
	byt :=encoder.Sum(nil)
	strr := hex.EncodeToString(byt)

	fmt.Println("aaa   "+encode)
	fmt.Println("bbb    "+strr)
}
