package models

import (
	"github.com/astaxie/beego/orm"
)

type Soarconfig struct {
	Id                  int
	Cfgname             string
	Cfgonline           string
	Cfgtest          	string
	Allowonlineastest	int8
	Sampling      		int8
}


func SoarconfigGetList(page, pageSize int, filters ...interface{}) ([]*Soarconfig, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Soarconfig, 0)
	query := orm.NewOrm().QueryTable("soarconfig")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)

	return list, total
}