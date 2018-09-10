package lot

import (
	"skinstore/common"
	"skinstore/utils/SqliteUtil"
)

type LotEntity struct {
	Pos string 		`json:"pos"`
	Status string 	`json:"status"`
}


func GetLot(pos string) LotEntity{
	db := SqliteUtil.NewSqlDb()
	var lot LotEntity
	stmt,err := db.Db.Prepare(`select pos as Pos,status as Status from lot_status where pos = ?`)
	common.CheckErr(err)
	defer stmt.Close()
	res:= stmt.QueryRow(pos)
	err = res.Scan(&lot.Pos,&lot.Status)
	return lot
}


func (lot *LotEntity) Update() bool{
	if lot == nil{
		return false
	}
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("update lot_status set status = ? where pos = ?")
	common.CheckErr(err)
	result,err := stmt.Exec(lot.Status,lot.Pos)
	common.CheckErr(err)
	affectNum ,err := result.RowsAffected()
	common.CheckErr(err)
	if affectNum > 0{
		return true
	}else{
		return false
	}
}


func (lot *LotEntity) Save() bool {
	if lot == nil {
		return false
	}
	db := SqliteUtil.NewSqlDb()
	stmt, err := db.Db.Prepare("insert into pos_status(`pos`,`status`) values(?,?)")
	common.CheckErr(err)
	_,err = stmt.Exec(lot.Pos,lot.Status)
	common.CheckErr(err)
	return true
}