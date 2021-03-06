package router

import (
	"net/http"
	"strings"
	"encoding/json"
	"skinstore/common"
	"errors"
	"fmt"
	"io/ioutil"
	"skinstore/Entity/wx/message"
	"encoding/xml"
	"skinstore/Entity/user"
	"skinstore/common/logger"
	"mime/multipart"
)

type Handler func(*Params,http.ResponseWriter)(*common.WebResult)
type TempHandler func(w http.ResponseWriter)
type CommHandler func(r *http.Request,w http.ResponseWriter)
var log = logger.NewLog()
var noLoginUrl = []string{"/api/wx/msg"}

var staticHander http.Handler = nil

type Router struct {
	Route map[string]map[string]Route
	Templ map[string]TempRoute
	Comm map[string]map[string]CommRoute
	prefix string
}

type Route struct {
	Handler Handler
	Params map[string]bool //是否必填参数
	Method string
	Path string
}

type TempRoute struct {
	Handler TempHandler
	Path string
}

type CommRoute struct{
	Path string
	Method string
	Handler CommHandler
}


func (r *Router)ServeHTTP(w http.ResponseWriter,req *http.Request){
	//优先处理静态文件
	if strings.HasPrefix(req.URL.Path,"/static/") && staticHander != nil{
		staticHander.ServeHTTP(w,req);
		return
	}

	//网页模板优先
	if route,ok:= r.Templ[req.URL.Path];ok{
		route.Handler(w)
		return
	}

	if route,ok := r.Comm[req.Method][req.URL.Path];ok{
		route.Handler(req,w);
		return
	}

	if route,ok := r.Route[req.Method][req.URL.Path];ok{
		params,err := getParams(req,route.Params)
		if err != nil {
			log.Error(err)
			http.Error(w,err.Error(),400)
		}else{
			res:=route.Handler(params,w)
			if res != nil{
				writeResp(w,res)
			}
		}
	}else{
		http.NotFound(w,req)
	}
}

func (r *Router)RegStatic(path string){
	staticHander = http.StripPrefix("/static/",http.FileServer(http.Dir(path)))
}

func (r *Router)RegHandlers(routes []Route){
	for _,route := range routes{
		method := strings.ToUpper(route.Method)
		if r.Route == nil {
			r.Route = make(map[string]map[string]Route)
		}
		if r.Route[method] == nil{
			r.Route[method] = make(map[string]Route)
		}
		r.Route[method][route.Path] = route
		log.Infof("http method:%s mappering for %s",route.Method,route.Path)
	}
}

func (r *Router)RegTemp(routes []TempRoute){
	if r.Templ == nil{
		r.Templ = make(map[string]TempRoute)
	}
	for _,route := range routes{
		r.Templ[route.Path] = route
		log.Infof("http template %s", route.Path)
	}
}

func (r *Router)RegComm(routes []CommRoute){
	for _,route := range routes{
		if(r.Comm == nil){
			r.Comm = make(map[string]map[string]CommRoute)
		}
		method := strings.ToUpper(route.Method)
		if r.Comm[method] == nil{
			r.Comm[method] = make(map[string]CommRoute)
		}
		r.Comm[route.Method][route.Path] = route
		log.Infof("http method:%s path:%s",route.Method,route.Path)
	}
}

type Params struct {
	data map[string]interface{}
}

func (p *Params) Get(key string)string{
	if v,ok := p.data[key];ok{
		return fmt.Sprintf("%v",v)
	}else{
		return ""
	}
}
func (p *Params) GetData()[]byte{
	if v,ok := p.data["json"];ok {
		return v.([]byte)
	}else{
		return nil
	}
}

func (p *Params) GetFile()multipart.File{
	if v,ok := p.data["file"];ok{
		return v.(multipart.File)
	}else{
		return nil
	}
}

func (p *Params) GetFileHeader() *multipart.FileHeader{
	if v,ok := p.data["header"];ok{
		return v.(*multipart.FileHeader)
	}else{
		return nil
	}
}


func (p *Params)GetWxMsg() *message.WxMsg{
	if v,ok := p.data["wx_msg_code"];ok{
		if msg,ok := v.(message.WxMsg);ok{
			return &msg
		}else{
			return nil
		}
	}else{
		return nil
	}
}

func getParams(req *http.Request,keys map[string]bool) (*Params,error){
	if keys != nil && len(keys) >0{
		if req.Method == "GET"{
			req.ParseForm()
			params := make(map[string]interface{})
			for key,isMust := range keys{
				value := req.Form.Get(key)
				if isMust {
					if value != ""{
						params[key] = value
					}else{
						return nil,errors.New(fmt.Sprintf("param:%s is not be nil",key))
					}
				}else{
					if value != ""{
						params[key] = value
					}
				}
			}
			return &Params{data:params},nil
		}else{
			//post
			paramMap := make(map[string]interface{})
			//if keys["upload"]{
			//	file,header,error := req.FormFile("file")
			//	if error != nil{
			//		paramMap["error"] = error
			//	}else{
			//		paramMap["file"] = file
			//		paramMap["header"] = header
			//	}
			//
			//	return &Params{data:paramMap},nil
			//}
			////get wx msg
			//if "/api/wx/msg" == req.URL.Path{
			//	paramMap["wx_msg_code"] = wxMsgParse(req)
			//	return &Params{data:paramMap},nil
			//}
			contentType := req.Header.Get("content-type")
			switch contentType {
			case "application/x-www-form-urlencoded":
				req.ParseForm()
				for k,v := range keys{
					value := req.PostForm.Get(k)
					if v {
						if value == "" {
							return  nil,errors.New(fmt.Sprintf("param:%s is not be nil",k))
						}
					}
					paramMap[k] = value
				}
				return &Params{data:paramMap},nil
			case "application/json":
				result, _:= ioutil.ReadAll(req.Body)
				json.Unmarshal(result,&paramMap)
				return &Params{data:paramMap},nil

			case "application/xml":
				paramMap["wx_msg_code"] = wxMsgParse(req)
				return &Params{data:paramMap},nil
			case "multipart/form-data":
				file,header,error := req.FormFile("file")
				if error != nil{
					paramMap["error"] = error
				}else{
					paramMap["file"] = file
					paramMap["header"] = header
				}
				return &Params{data:paramMap},nil
			default:
				return &Params{data:paramMap},errors.New("do not suppor this request")
			}
		}
	}else{
		return nil,nil
	}
}

func wxMsgParse(r *http.Request) message.WxMsg{
	data,_ := ioutil.ReadAll(r.Body)
	var msg message.WxMsg
	err := xml.Unmarshal(data,&msg)
	common.CheckErr(err)
	return msg
}


func writeResp(w http.ResponseWriter,res *common.WebResult){
	w.Header().Set("Content-Type","application/json")
	b,err := json.Marshal(res)
	common.CheckErr(err)
	w.Write(b)
}

/**
检查用户是否登录了
 */
func authCheck(req *http.Request) bool {
	for _,v := range noLoginUrl{
		if v == req.URL.Path{
			return true
		}
	}
	userIdCookie,err := req.Cookie("userId")
	common.CheckErr(err)
	openIdCookie,err := req.Cookie("openId")
	common.CheckErr(err)
	return user.IsLoginUser(userIdCookie.Value,openIdCookie.Value)

}