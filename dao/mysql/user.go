package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go_web_app/models"

	"go.uber.org/zap"
)

const serect = "wuyue.com"

// 定义 error的常量方便判断
var (
	UserAleadyExists = errors.New("用户已存在")
	WrongPassword    = errors.New("密码不正确")
	UserNoExists     = errors.New("用户不存在")
)

// dao层 其实就是将数据库操作 封装为函数 等待logic层 去调用她

func InsertUser(user *models.User) error {
	// 密码要加密保存
	user.Password = encryptPassword(user.Password)
	sqlstr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err := db.Exec(sqlstr, user.UserId, user.Username, user.Password)
	if err != nil {
		zap.L().Error("InsertUser dn error", zap.Error(err))
		return err
	}
	return nil
}

//
func Login(user *models.User) error {
	oldPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err := db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return UserNoExists
	}
	if err != nil {
		return err
	}
	if encryptPassword(oldPassword) != user.Password {
		return WrongPassword
	}
	return nil
}

// CheckUserExist 检查数据库是否有该用户名
func CheckUserExist(username string) error {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlstr, username)
	if err != nil {
		zap.L().Error("CheckUserExist dn error", zap.Error(err))
		return err
	}
	if count > 0 {
		return UserAleadyExists
	}
	return nil
}

// 加密密码
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(serect))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
