package controllers

//c.JSON(http.StatusOK, &Response{
//Code: CodeSuccess,
//Msg:  CodeSuccess.Msg(),
//Data: data,
//})

type _ResponseLogin struct {
	Code int64 `json:"code"` // 业务状态响应码
	Message string `json:"message"` //提示信息
	Data  string `json:"data"` // token
}

type _RequestLogin struct {
	// 用户名
	Username string `json:"username"`
	//密码
	Password string `json:"password"`
}