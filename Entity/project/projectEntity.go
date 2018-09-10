package project

import (
	"skinstore/utils/SqliteUtil"
	"skinstore/common"
	"fmt"
	"strings"
	"errors"
)

type ProjectEntity struct{
	Id int  				`json:"id"`
	Name string				`json:"name"`
	Description string		`json:"description"`
	Type string				`json:"type"`
	ImgUrl string			`json:"imgurl"`
	OriginalPrice int		`json:"originalPrice"`
	CurPrice int			`json:"curPrice"`
	Status byte				`json:"status"`
}

func GetAllProjectList(index,rows int) []ProjectEntity{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare(`select id as Id,name as Name, description as Description,
				type as Type,img_url as ImgUrl,original_price as OriginalPrice,cur_price as CurPirce,
				status as Status from project limit ?,?`)
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Query(index,rows)
	common.CheckErr(err)
	var list []ProjectEntity
	for res.Next(){
		var project ProjectEntity
		err := res.Scan(&project.Id,&project.Name,&project.Description,&project.Type,&project.ImgUrl,&project.OriginalPrice,&project.CurPrice,&project.Status)
		common.CheckErr(err)
		list = append(list,project)
	}
	return list
}

func GetProjectListByParam(params map[string]string) []ProjectEntity{
	db := SqliteUtil.NewSqlDb()
	sql := []string{"select id as Id,name as Name, description as Description,type as Type,img_url as ImgUrl,original_price as OriginalPrice,cur_price as CurPirce, status as Status from project where status=1"}
	if paramType,ok := params["type"];ok{
		sql = append(sql,fmt.Sprintf(" and type='%s'",paramType))
	}
	index,_ := params["index"]
	rows,_ := params["rows"]
	sql = append(sql,fmt.Sprintf(" order by id desc limit %s,%s",index,rows))
	stmt,err := db.Db.Prepare(strings.Join(sql,""))
	fmt.Println(strings.Join(sql,""))
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Query()
	common.CheckErr(err)
	var list []ProjectEntity
	for res.Next(){
		var project ProjectEntity
		err := res.Scan(&project.Id,&project.Name,&project.Description,&project.Type,&project.ImgUrl,&project.OriginalPrice,&project.CurPrice,&project.Status)
		common.CheckErr(err)
		list = append(list,project)
	}
	return list
}

/**
新增项目
 */
func (p *ProjectEntity)Save() error{
	db := SqliteUtil.NewSqlDb()
	stmt,error := db.Db.Prepare("insert into project(`name`,`description`,`type`,`img_url`,`original_price`,`cur_price`,`status`)values(?,?,?,?,?,?,?)")
	common.CheckErr(error)
	defer stmt.Close()
	res,error := stmt.Exec(p.Name,p.Description,p.Type,p.ImgUrl,p.OriginalPrice,p.CurPrice,p.Status)
	common.CheckErr(error)
	if rows,err := res.RowsAffected();rows < 1 || err != nil{
		return errors.New("save project to sql fail")
	}else{
		return nil
	}
}
