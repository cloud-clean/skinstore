package TempParser

import (
	"html/template"
	"net/http"
)

func LotLogin(rw http.ResponseWriter){
	t,_:=template.ParseFiles("template/lotLogin.html")
	t.Execute(rw,"")
}
