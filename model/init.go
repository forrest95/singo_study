package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"singo/util"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

var Sqlx_db *sqlx.DB //fang自添加sqlx功能

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

//初始化sqlx //fang 自定义
func InitDB_sqlx(connString string) (err error) {
	fmt.Println("初始化了InitDB_sqlx")
	dsn := connString
	// 也可以使用MustConnect连接不成功就panic
	Sqlx_db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	Sqlx_db.SetMaxOpenConns(20)
	Sqlx_db.SetMaxIdleConns(10)

	return
}
