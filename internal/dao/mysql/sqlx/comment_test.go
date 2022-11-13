package sqlx

import (
	"fmt"
	"testing"
	"wenba/internal/dao/mysql"
	"wenba/internal/model/config"
)

//5688866677723136

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
func TestGetCommentByID(t *testing.T) {
	commentID := 5688866677723136
	commentInfo, _ := GetCommentByID(int64(commentID))
	fmt.Println("commentInfo:", commentInfo)
	fmt.Println("authorID", commentInfo.AuthorID)
}
