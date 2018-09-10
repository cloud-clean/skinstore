package Handler

import (
	"net/http"
	"skinstore/common"
	"skinstore/Entity/project"
	"strconv"
	"skinstore/web/router"

	"skinstore/common/logger"
	"encoding/json"
)

var log = logger.NewLog()

/**
项目列表
 */
func ProjectListHander(params *router.Params,rw http.ResponseWriter) *common.WebResult{
		rowsStr := params.Get("rows")
		rows,err := strconv.Atoi(rowsStr)
		if err != nil {
			//log.Error(err)
			rows = 15
		}
		pageStr := params.Get("page")
		page,err := strconv.Atoi(pageStr)
		if err != nil{
			//log.Error(err)
			page = 0
		}
		index := (page-1)*rows
		list := project.GetAllProjectList(index,rows)
		return common.NewResult(1,list)
}

/**
根据项目类型返回项目列表
 */
func ProjectLisByTypetHander(p *router.Params,rw http.ResponseWriter)*common.WebResult{
	params := make(map[string]string)
	rowsStr := p.Get("rows")
	if rowsStr == ""{
		rowsStr = strconv.Itoa(common.ROWS_SIZE)
	}
	params["rows"] = rowsStr
	indexStr := p.Get("index")
	if indexStr == ""{
		indexStr = "0"
	}
	params["index"] = indexStr
	if typep := p.Get("type");typep != ""{
		params["type"] = typep
	}
	list := project.GetProjectListByParam(params)
	return common.NewResult(1,list)
}
/**
添加项目
 */
func ProjectAddHandler(p *router.Params,rw http.ResponseWriter)*common.WebResult{
	var entity project.ProjectEntity
	error := json.Unmarshal(p.GetData(),&entity)
	if error != nil{
		return common.NewResult(0,error.Error())
	}
	error = entity.Save()
	if error != nil{
		return common.NewResult(0,error.Error())
	}
	return common.NewResult(1,nil)
}

