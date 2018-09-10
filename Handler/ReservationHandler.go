package Handler

import (
	"net/http"
	"strconv"
	"skinstore/common"
	"skinstore/Entity/reservation"
	"time"
	"skinstore/web/router"
	"skinstore/common/json"
)

func AddReservationHandler(params *router.Params,rw http.ResponseWriter)*common.WebResult{
	uid := params.Get("uid")
	pId := params.Get("projectId")
	reservTm := params.Get("reservTm")
	mobile := params.Get("mobile")
	name := params.Get("name")
	timestamp,err := strconv.ParseInt(reservTm,10,64)
	rserTm := json.JsonTime{Data:time.Unix(timestamp,0)}
	common.CheckErr(err)
	userId,err := strconv.Atoi(uid)
	common.CheckErr(err)
	projectId,err := strconv.Atoi(pId)
	common.CheckErr(err)
	entity := &reservation.ReservEntity{UserId:userId,ProjectId:projectId,Mobile:mobile,Name:name,ReservTm:rserTm}
	err = entity.Save()
	if err != nil{
		return common.NewResult(common.FAIL,err.Error())
	}
	return common.NewResult(common.SUCCESS,"success")
}

func GetAllReservationHandler(params *router.Params,rw http.ResponseWriter)*common.WebResult{
	rowsStr := params.Get("rows")
	rows,err := strconv.Atoi(rowsStr)
	if err != nil {
		log.Error(err)
		rows = 15
	}
	pageStr := params.Get("page")
	page,err := strconv.Atoi(pageStr)
	if err != nil{
		log.Error(err)
		page = 0
	}
	statusStr := params.Get("status")
	var status int
	var startTm time.Time
	if statusStr == ""{
		status = -1
	}else{
		status,err = strconv.Atoi(statusStr)
		common.CheckErr(err)
	}
	timeStr := params.Get("startTm")
	if timeStr == ""{
		startTm = time.Unix(0,0)
	}else{
		ts,err := strconv.ParseInt(timeStr,10,64)
		common.CheckErr(err)
		startTm = time.Unix(ts,0)
	}
	index := (page-1)*rows
	list,total := reservation.GetReservationList(status,startTm,index,rows)
	res := common.NewResult(common.SUCCESS,list)
	res.SetData("total",total)
	return res
}

func GetTodayReservationHandler(params *router.Params,rw http.ResponseWriter)*common.WebResult{
	rowsStr := params.Get("rows")
	rows,err := strconv.Atoi(rowsStr)
	if err != nil {
		log.Error(err)
		rows = 15
	}
	pageStr := params.Get("page")
	page,err := strconv.Atoi(pageStr)
	if err != nil{
		log.Error(err)
		page = 0
	}
	index := (page-1)*rows
	list := reservation.GetTodayReservList(index,rows)
	return common.NewResult(common.SUCCESS,list)
}

func UpdateReservationStatusHandler(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	statusStr := params.Get("status")
	status,err := strconv.Atoi(statusStr)
	common.CheckErr(err)
	id := params.Get("id")
	idint,err := strconv.Atoi(id)
	common.CheckErr(err)
	var entity = reservation.ReservEntity{Id:idint,Status:status}
	if entity.UpdateStatus(){
		return common.NewResult(common.SUCCESS,"success")
	}
	return common.NewResult(common.FAIL,"failed")
}
