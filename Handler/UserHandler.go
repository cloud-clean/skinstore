package Handler

import (
	"skinstore/web/router"
	"net/http"
	"skinstore/common"
)

func UserHandler(params *router.Params,rw http.ResponseWriter)*common.WebResult{
	return common.NewResult(2,"")
}
