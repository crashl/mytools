package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mytools/models"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.Layout = "layout.html"
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	//1.拿到post数据
	username := c.GetString("username")
	password := c.GetString("password")
	beego.Info(username, password)
	//2.判断数据合法性
	if username == "" || password == "" {
		beego.Info("注册信息不完整")
		c.Redirect("/register", 302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	t_user := models.User{}
	t_user.Name = username
	t_user.Pwd = password
	_, err := o.Insert(&t_user)
	if err != nil {
		beego.Info("插入数据库失败")
		c.Redirect("/register", 302)
		return
	}
	//4.返回登录界面
	c.Layout = "layout.html"
	c.TplName = "login.html"
}
