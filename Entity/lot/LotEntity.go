package lot

import (
	"skinstore/common"
	"skinstore/utils/SqliteUtil"
)

type LampStatusEntity struct {
	Pos string 		`json:"pos"`
	Status string 	`json:"status"`
}


func GetLot(pos string) *LampStatusEntity{
	db := SqliteUtil.NewSqlDb()
	var lot LampStatusEntity
	stmt,err := db.Db.Prepare(`select pos as Pos,status as Status from lot_status where pos = ?`)
	common.CheckErr(err)
	defer stmt.Close()
	res:= stmt.QueryRow(pos)
	err = res.Scan(&lot.Pos,&lot.Status)
	if err != nil{
		return nil
	}
	return &lot
}


func (lot *LampStatusEntity) Update() bool{
	if lot == nil{
		return false
	}
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("insert or replace into lot_status(pos,status) values(?,?)")
	common.CheckErr(err)
	result,err := stmt.Exec(lot.Pos,lot.Status)
	common.CheckErr(err)
	affectNum ,err := result.RowsAffected()
	common.CheckErr(err)
	if affectNum > 0{
		return true
	}else{
		return false
	}
}


func (lot *LampStatusEntity) Save() bool {
	if lot == nil {
		return false
	}
	db := SqliteUtil.NewSqlDb()
	stmt, err := db.Db.Prepare("insert into lot_status(`pos`,`status`) values(?,?)")
	common.CheckErr(err)
	_,err = stmt.Exec(lot.Pos,lot.Status)
	common.CheckErr(err)
	return true
}