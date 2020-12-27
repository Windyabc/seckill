package access

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "seckill-backend/models/access"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

// 注册展示页面
func (this *RegisterController) ShowRegister() {
	this.TplName = "register.html"
}

// 注册获取数据页面
func (this *RegisterController) HandleRegister() {
	// 获取浏览器传递的值，并去除两边的空格

	userName := strings.TrimSpace(this.GetString("userName"))
	passWord := strings.TrimSpace(this.GetString("passWord"))
	if userName == "" || passWord == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "register.html"
		this.Data["errmsg"] = "用户名或密码不能为空"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象x
	o := orm.NewOrm()
	//   获取插入对象
	user := models.User{}
	//   插入数值
	user.UserName = userName
	user.PassWord = passWord

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
		this.TplName = "register.html"
		this.Data["errmsg"] = "插入数据失败"
		return
	}
	// 测试返回视图
	this.Ctx.WriteString("注册成功！")
	// 实际情况注册成功返回到登录页面
	// this.Redirect("login", 302)
}