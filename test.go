package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "cmd:haha"
	fmt.Println(str[:strings.Index(str,":")])
	fmt.Println(str[strings.Index(str,":")+1:])

}
