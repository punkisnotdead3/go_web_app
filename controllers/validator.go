package controllers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

// 全局翻译器
var trans ut.Translator

// InitTrans locale 指定你想要的翻译 环境
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//注册一个获取jsonTag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		//中文
		zhT := zh.New()
		//英文
		enT := en.New()
		// 第一个参数 是备用的语言
		uni := ut.New(enT, zhT, enT)

		//local 一般会在前端的请求头中 定义Accept-Language
		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator failed:%s ", locale)
		}

		switch locale {
		case "en":
			err = enTrans.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTrans.RegisterDefaultTranslations(v, trans)
		default:
			// 默认是英文
			err = enTrans.RegisterDefaultTranslations(v, trans)
		}
		return err
	}
	return err
}

// 去除报错信息中的结构体信息
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		// 这里 算法非常简单 就是遍历你的错误信息 然后把key值取出来 把.之前的信息去掉就行了
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
