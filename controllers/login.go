package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mytools/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Layout = "layout.html"
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	//1.拿到post数据
	username := c.GetString("username")
	password := c.GetString("password")
	//2.判断数据合法性
	if username == "" || password == "" {
		beego.Info("注册信息不完整")
		c.Redirect("/register", 302)
		return
	}
	//3.查询账号密码是否正确
	o := orm.NewOrm()
	t_user := models.User{}
	t_user.Name = username
	t_user.Pwd = password
	err := o.Read(&t_user, "Name")
	if err != nil {
		beego.Info("没有用户信息")
		c.Layout = "layout.html"
		c.TplName = "login.html"
		return
	}
	//4.跳转页面
	c.Ctx.WriteString("欢迎登录")
}
