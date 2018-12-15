package utils

import (
	"math/rand"
	"time"
)

func genRand(len int)string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte,len)
	for i:= 0;i<len;i++{
		b:=r.Intn(52)+65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
