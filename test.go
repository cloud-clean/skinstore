package main

import (
	"fmt"
	"skinstore/tcp"
	"strings"
)

func main(){
	str := "cmd:haha"
	fmt.Println(str[:strings.Index(str,":")])
	fmt.Println(str[strings.Index(str,":")+1:])
	tcp.Start()
	//var flag = make(chan int,1)
	//flag <- 1
	//fmt.Println(len(flag))

}
