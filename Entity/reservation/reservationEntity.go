package reservation

import (
	"skinstore/utils/SqliteUtil"
	"skinstore/common"
	"errors"
	"skinstore/common/json"
	"time"
	"strings"
	sql2 "database/sql"
)

type ReservEntity struct {
	Id int					`json:"id"`
	UserId int				`json:"userId"`
	ProjectId int			`json:"projectId"`
	ReservTm json.JsonTime	`json:"reservTm"`
	Mobile string			`json:"mobile"`
	Name string				`json:"name"`
	Status int				`json:"status"`
}

func (r *ReservEntity) Save() error{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare("insert into reservation(user_id,name,project_id,mobile,re_time) values(?,?,?,?,?);")
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Exec(r.UserId,r.Name,r.ProjectId,r.Mobile,r.ReservTm.Data)
	common.CheckErr(err)
	if rows,err := res.RowsAffected();rows < 1 || err != nil{
		return errors.New("save reservation to sql fail")
	}
	return nil
}

func GetReservationList(status int,startTm time.Time,index,rows int) ([]ReservEntity,int){
	db := SqliteUtil.NewSqlDb()
	sql := []string{`select id as Id,name as Name,user_id as UserId,mobile as Mobile,project_id as ProjectId,re_time as ReservTm,status as Status from reservation `}
	countSql := []string{`select count(0) count from reservation `}
	if status > -1 {
		sql = append(sql,"where status = ?")
		countSql = append(countSql,"where status = ?")
	}else{
		sql = append(sql,"where status > ?")
		countSql = append(countSql,"where status > ?")
	}
	sql = append(sql," and re_time > ?")
	countSql = append(countSql, " and re_time > ?")
	sql = append(sql," order by create_tm limit ?,?")
	stmt,err := db.Db.Prepare(strings.Join(sql,""))
	countStmt,err := db.Db.Prepare(strings.Join(countSql,""))
	common.CheckErr(err)
	var res *sql2.Rows
	var countRes *sql2.Rows
	res,err = stmt.Query(status,startTm,index,rows)
	countRes,err = countStmt.Query(status,startTm)
	defer stmt.Close()
	defer countStmt.Close()
	common.CheckErr(err)
	var list []ReservEntity
	for res.Next(){
		var entity ReservEntity
		err = res.Scan(&entity.Id,&entity.Name,&entity.UserId,&entity.Mobile,&entity.ProjectId,&entity.ReservTm.Data,&entity.Status)
		common.CheckErr(err)
		list = append(list,entity)
	}
	var total = 0
	for countRes.Next(){
		err = countRes.Scan(&total)
		common.CheckErr(err)
	}
	return list,total
}
/**
修改预约状态
 */
func (e *ReservEntity)UpdateStatus() bool{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare(`update reservation set status = ? where id = ?`)
	common.CheckErr(err)
	res,err := stmt.Exec(e.Status,e.Id)
	common.CheckErr(err)
	num,_ := res.RowsAffected()
	if num >0{
		return true
	}else{
		return false
	}
}

/**
查看当天预约列表
 */

 func GetTodayReservList(index,rows int) []ReservEntity{
 	db := SqliteUtil.NewSqlDb()
 	res,err := db.Db.Query(`select id as Id,name as Name,user_id as UserId,mobile as Mobile,project_id as ProjectId,re_time as ReservTm,status as Status 
				from reservation where status = 0 and re_time 
				between datetime('now','start of day','+0 second') and datetime('now','start of day','+1 day') limit ?,?`,index,rows)
 	common.CheckErr(err)
	 var list []ReservEntity
	 for res.Next(){
		 var entity ReservEntity
		 err = res.Scan(&entity.Id,&entity.Name,&entity.UserId,&entity.Mobile,&entity.ProjectId,&entity.ReservTm.Data,&entity.Status)
		 common.CheckErr(err)
		 list = append(list,entity)
	 }
	 return list
 }