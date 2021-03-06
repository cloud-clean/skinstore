package main

import (
	"fmt"
	"net/http"
	"skinstore/common/logger"
	"skinstore/mqttClient"
	"skinstore/web"
	"skinstore/web/router"
	"strconv"
	"time"
)

func main(){
	port := 8082

	//http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("../template/static"))))

	r := &router.Router{}
	r.RegHandlers(web.InitRoute())
	r.RegTemp(web.InitTemplate())
	r.RegComm(web.InitComm())
	r.RegStatic("./template/static")
	svr :=http.Server{
		Addr:":"+strconv.Itoa(port),
		ReadTimeout:time.Second*5,
		WriteTimeout:time.Second*5,
		Handler:r,
	}


	mqttClient.MqttInit()
	//mqttClient.Mc.Subscribe(func(client mqtt.Client, message mqtt.Message) {
	//	fmt.Println("get msg form mqtt:"+string(message.Payload()))
	//})
	logger.NewLog().Infof("start server listen on:%s",svr.Addr)
	err := svr.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}
}

