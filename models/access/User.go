package access

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户信息
type User struct {
	Id       int `orm:"column(id);auto`
	UserName string `orm:"column(username);unique` // 用户名唯一
	PassWord string `orm:"column(password)`
}

func init() {
	// 映射modle数据
	orm.RegisterModel(new(User))
}
