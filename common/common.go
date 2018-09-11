package common

import (
	"strconv"
)

const(
	ROWS_SIZE = 15
)

type WebResult struct {
	Code int    `json:"code"`
	Msg string 	`json:"msg"`
	Result map[string]interface{}		`json:"result"`
}

func NewResult(code int,res interface{}) *WebResult{
	msg := NewConfig().Get("errcode",strconv.Itoa(code))
	data := map[string]interface{}{
		"data":res,
	}
	return &WebResult{Code:code,Msg:msg,Result:data,}
}

func (res *WebResult)SetData(key string,data interface{}){
	res.Result[key]=data
}

func CheckErr(err error){
	if err != nil {
		panic(err)
	}
}


