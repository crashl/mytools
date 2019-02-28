package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beegodemo/models"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
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
	c.TplName = "login.html"
}

//登录也做在这个文件，是为了验证路由用其他函数替代get和post
func (c *RegisterController) ShowLogin(){
	c.TplName = "login.html"
}

func (c *RegisterController) HandleLogin(){
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
	err := o.Read(&t_user,"Name")
	if err != nil{
		beego.Info("没有用户信息")
		c.TplName = "login.html"
		return
	}
	//4.跳转页面
	c.Ctx.WriteString("欢迎登录")
}