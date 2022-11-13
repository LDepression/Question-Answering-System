package sqlx

import (
	"wenba/internal/dao/mysql"
	"wenba/internal/model/config"
)

//在测试之前,我们需要对mysql进行init
func init() {
	dbCfg := config.MySQL{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "zxz123456",
		DbName:       "wenba",
		Port:         "3306",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	mysql.Init(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DbName)
}

//func TestInsertUser(t *testing.T) {
//	user := &User{
//		ID:        1,
//		UserName:  "yyy",
//		Password:  "123",
//		Email:     "1231@qq.com",
//		Privilege: "用户",
//		Gender:    "男",
//	}
//	InsertUser(user)
//}
