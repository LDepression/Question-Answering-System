package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"wenba/internal/global"
)

var Db *sqlx.DB

func Init(user, password, host, port, dbname string) *sqlx.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, dbname)
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.SetMaxIdleConns(global.Settings.MySQL.MaxIdleConns)
	Db.SetMaxOpenConns(global.Settings.MySQL.MaxOpenConns)
	return Db
	////初始化连接
	//	func initDB() (err error) {
	//	//数据库连接信息：dataSourceName
	//	dsn := "root:123123@tcp(192.168.241.129:3306)/go_demo"
	//	db, err = sqlx.Connect("mysql", dsn) //判断dsn格式即账号密码是否正确
	//	if err != nil {
	//		return
	//	}
	//	//db.SetMaxOpenConns(10) //设置连接池中最大的连接数
	//	//db.SetMaxIdleConns(5)  //设置连接池中的最大闲置连接数
	//	return
}
