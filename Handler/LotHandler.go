package Handler

import (
	"net/http"
	"skinstore/Entity/lot"
	"skinstore/common"
	"skinstore/web/router"
)

func LampStatusHander(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	pos := params.Get("pos")
	if pos != ""{
		lotEntity := lot.GetLot(pos)
		if lotEntity == nil{
			return common.NewResult(0,"pos is not exits")
		}
		return common.NewResult(1,lotEntity)
	}else{
		return common.NewResult(0,"pos is null")
	}
}


func SaveLampHander(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	pos := params.Get("pos")
	status := params.Get("status")

	if pos != "" && status != ""{
		var lot = lot.LotEntity{Pos:pos,Status:status}
		if lot.Save(){
			return common.NewResult(1,"success")
		}
		return common.NewResult(0,"false")
	}else{
		return common.NewResult(0,"pos is null")
	}
}



func UpdateLampHander(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	pos := params.Get("pos")
	status := params.Get("status")

	if pos != "" && status != ""{
		var lot = lot.LotEntity{Pos:pos,Status:status}
		if lot.Update(){
			return common.NewResult(1,"success")
		}
		return common.NewResult(0,"false")
	}else{
		return common.NewResult(0,"pos is null")
	}
}