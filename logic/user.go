package logic

//logic 其实就是存放业务层的代码

import (
	"go_web_app/dao/mysql"
	"go_web_app/models"
	"go_web_app/pkg/snowflake"
)

func Login(login *models.ParamLogin) error {
	user := models.User{
		Username: login.UserName,
		Password: login.Password,
	}
	return mysql.Login(&user)
}

func Register(register *models.ParamRegister) (err error) {
	// 判断用户是否存在
	err = mysql.CheckUserExist(register.UserName)
	if err != nil {
		// db 出错
		return err
	}
	// 生成userid
	userId := snowflake.GenId()
	// 构造一个User db对象
	user := models.User{
		UserId:   userId,
		Username: register.UserName,
		Password: register.Password,
	}
	// 保存数据库
	err = mysql.InsertUser(&user)
	if err != nil {
		return err
	}
	return
}
