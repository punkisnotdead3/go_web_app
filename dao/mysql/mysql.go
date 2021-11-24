package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Close() {
	_ = db.Close()
}

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"), viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_connection"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_connection"))
	return
}
