package Handler

import (
	"net/http"
	"html/template"
	"skinstore/web/router"
	"skinstore/common"
	"os"
	"io"
)

var FILE_PATH = "E:/upload/"


func UploadHtmlHandler(p *router.Params,w http.ResponseWriter) *common.WebResult{
	t,err := template.ParseFiles("template/upload.html")
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return nil
	}
	t.Execute(w,nil)
	return nil
}

func UploadFileHandler(p *router.Params,w http.ResponseWriter) *common.WebResult{
	file := p.GetFile()
	header := p.GetFileHeader()
	if file != nil  && header != nil{
		defer file.Close()
		fw,err := os.Create(FILE_PATH+header.Filename)
		if err != nil{
			return common.NewResult(0,err)
		}
		defer fw.Close()
		_,err = io.Copy(fw,file)
		if err != nil{
			return common.NewResult(0,nil)
		}
		return common.NewResult(1,"success")
	}else{
		return common.NewResult(0,"file is nil")
	}
}
