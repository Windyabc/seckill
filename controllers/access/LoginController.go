package access

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "seckill-backend/models/access"
	"strings"
)

type LoginController struct {
	beego.Controller
}

// 登录页面 get
func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}

// 登录页面 post
func (this *LoginController) HandleLogin() {
	// 拿到浏览器数据，并去除两边的空格
	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	beego.Info("账号:", Name, "密码:", Pwd)

	//数据处理
	if Name == "" || Pwd == "" {
		beego.Info("登录失败！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "登录失败！"
		return
	}
	// 查找数据库
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var user models.User
	// 查询
	user.UserName = Name
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("用户名登录失败！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "用户名登录失败！"
		return
	}
	// 判断密码是否一致
	if user.PassWord != Pwd {
		beego.Info("密码登录失败！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "密码登录失败！"
		return
	}
	// 测试返回视图
	this.Ctx.WriteString("登录成功")
}
