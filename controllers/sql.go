package controllers

import (
	"github.com/astaxie/beego"
	"mytools/models"
	"strings"
)

type SqlController struct {
	beego.Controller
	pagesize int
}

func (c *SqlController) Get() {
	c.Layout = "layout.html"
	c.TplName = "sql.html"
	c.Table()
}

func (c *SqlController) Post() {
	c.Layout = "layout.html"
	c.TplName = "sql.html"
}


func (self *SqlController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}
	self.pagesize = limit
	id := strings.TrimSpace(self.GetString("id"))
	//查询条件
	filters := make([]interface{}, 0)
	if id != "" {
		filters = append(filters, "id", id)
	}

	result, count := models.SoarconfigGetList(page, self.pagesize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["cfgname"] = v.Cfgname
		row["cfgonline"] = v.Cfgonline
		row["cfgtest"] = v.Cfgtest
		row["allowonlineastest"] = v.Allowonlineastest
		row["sampling"] = v.Sampling
		list[k] = row
	}

	beego.Info(count)
}