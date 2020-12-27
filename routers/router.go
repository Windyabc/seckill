package routers

import (
	"github.com/astaxie/beego"
	accessControllers "seckill-backend/controllers/access"
)

func init() {
	beego.Router("/", &accessControllers.IndexController{}, "get:ShowIndex")
	beego.Router("/register", &accessControllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &accessControllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
}
