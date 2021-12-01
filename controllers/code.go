package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeInvalidPassword
	CodeServerBusy
	CodeTokenIsEmpty
	CodeTokenInvalid
	CodeNoLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeInvalidPassword: "用户名或密码不正确",
	CodeServerBusy:      "服务繁忙 请稍后再试",
	CodeTokenIsEmpty:    "token 为空",
	CodeTokenInvalid:    "token非法",
	CodeNoLogin:         "未登录",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
