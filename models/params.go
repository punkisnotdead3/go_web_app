package models

//定义请求的参数结构体

type ParamRegister struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// eqfield 指定必须相等的字段
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
