package logic

//logic 其实就是存放业务层的代码

import (
	"go_web_app/dao/mysql"
	"go_web_app/models"
)

func Register(register *models.ParamRegister) {
	// 判断用户是否存在
	mysql.QueryUserByUserName()
	// 保存数据库
	//userId := snowflake.GenId()

	mysql.InsertUser()
}
