package main

import (
	"fmt"
	"skinstore/Entity/lot"
)

func main(){
	//str := "cmd:haha"
	//fmt.Println(str[:strings.Index(str,":")])
	//fmt.Println(str[strings.Index(str,":")+1:])
	//tcp.Start()
	user := lot.LotUser{Account:"test",Passworld:"teset"}
	err := user.Save();
	if(err != nil){
		fmt.Println(err.Error())
	}
}
