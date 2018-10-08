package web

import (
	"skinstore/CommHandler"
	"skinstore/TempParser"
	"skinstore/web/router"
	"skinstore/common"
	"net/http"
	"skinstore/Handler"
	"skinstore/common/logger"
)
var log = logger.NewLog()
var routers = []router.Route{}
var tempRouters = []router.TempRoute{}
var commRouters = []router.CommRoute{}
func InitRoute() []router.Route{
	//test
	addGet("/test/go",test, map[string]bool{"name":true,"age":false})
	addPost("/test/go",posttest, map[string]bool{"name":true,"age":false})
	//project
	addGet("/api/project/list",Handler.ProjectListHander,
		map[string]bool{"page":false,"rows":false})
	addGet("/api/project/list/type",Handler.ProjectLisByTypetHander,
		map[string]bool{"page":false,"rows":false,"type":false})
	addPost("/api/project/add",Handler.ProjectAddHandler,map[string]bool{"json":true})
	//reservation
	addPost("/api/reser/add",Handler.AddReservationHandler,map[string]bool{"uid":true,"projectId":true,"mobile":true,"reservTm":true,"name":true})
	addGet("/api/reser/list",Handler.GetAllReservationHandler,map[string]bool{"status":false,"startTm":false,"page":false,"rows":false})
	addGet("/api/reser/list/today",Handler.GetTodayReservationHandler,map[string]bool{"page":false,"rows":false})
	addPost("/api/reser/status/update",Handler.UpdateReservationStatusHandler,map[string]bool{"id":true,"status":true})
	//upload
	addPost("/api/upload",Handler.UploadFileHandler,map[string]bool{"upload":true})

	//lot
	addPost("/api/lot/update",Handler.UpdateLampHander,map[string]bool{"pos":true,"status":true})


	//weixin
	addGet("/api/wx/msg",Handler.MsgGetHandler,map[string]bool{"signature":true,"timestamp":true,"nonce":true,"echostr":true})
	addPost("/api/wx/msg",Handler.MsgPostHandler,map[string]bool{"data":true})
	addGet("/api/wx/oauth",Handler.WxOauthHandler,map[string]bool{"code":true,"state":true})
	addGet("/api/wx/sign",Handler.JsapiSignHandler,map[string]bool{"url":true})

	addGet("/api/upload",Handler.UploadHtmlHandler,map[string]bool{})

	//lot
	addGet("/api/lot",Handler.LampStatusHander,map[string]bool{"pos":true})
	return routers
}

func InitTemplate()[]router.TempRoute{
	addTemplate("/lot/login",TempParser.LotLogin)
	return tempRouters
}

func InitComm()[]router.CommRoute{
	addCommRoute("POST","/api/lot/login",CommHandler.LotLoginHandler)
	return commRouters;
}







func addGet(path string,h router.Handler,params map[string]bool){
	routers = append(routers,router.Route{Method:"get",Path:path,Handler:h,Params:params})
}

func addPost(path string,h router.Handler,params map[string]bool){
	routers = append(routers,router.Route{Method:"post",Path:path,Handler:h,Params:params})
}

func addTemplate(path string,handler router.TempHandler){
	tempRouters = append(tempRouters,router.TempRoute{Path:path,Handler:handler})
}

func addCommRoute(method string,path string,handler router.CommHandler){
	commRouters = append(commRouters,router.CommRoute{Handler:handler,Path:path,Method:method})
}


func test(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	log.Infof("name:%s",params.Get("name"))
	log.Infof("age:%s",params.Get("age"))
	return common.NewResult(1,"good")
}



func posttest(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	log.Infof("name:%s",params.Get("name"))
	log.Infof("age:%s",params.Get("age"))
	return common.NewResult(1,"good")
}