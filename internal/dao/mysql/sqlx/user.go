package sqlx

import (
	"errors"
	"fmt"
	"wenba/internal/dao/mysql"
	"wenba/internal/model/request"
	"wenba/internal/pkg/utils"
)

func InsertUser(user User) error {
	sqlStr := `insert into user(user_id,username,password,email,privilege,gender,avatar)
		values(?,?,?,?,?,?,?)`
	_, err := mysql.Db.Exec(sqlStr, user.ID, user.UserName, user.Password, user.Email, user.Privilege, user.Gender, user.Avatar)
	if err != nil {
		return err
	}
	return nil
}

func CheckUsername(username string) error {
	var userId int64
	sqlStr := `select user_id
	from user
	where username=?
`
	err := mysql.Db.Get(&userId, sqlStr, username)
	fmt.Println(err)
	return err
}

func CheckUsernameAndPassword(ParamLogin request.ReqLogin) error {

	var userID string
	sqlStr := `select user_id
	from user
	where username=? and password=?
`
	fmt.Println(ParamLogin.Username)
	fmt.Println(ParamLogin.Password)
	if err := mysql.Db.Get(&userID, sqlStr, ParamLogin.Username, ParamLogin.Password); err != nil {
		return err
	}
	fmt.Println("userID:", userID)
	if utils.StringToIDMust(userID) <= 0 {
		return errors.New("账号或密码错误")
	}
	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user = new(User)
	sqlStr := `
	select user_id,userName,password,email,privilege,gender,avatar
	from user
	where username =?
`
	if err := mysql.Db.Get(user, sqlStr, username); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(userID int64) (*User, error) {
	var user = new(User)
	sqlStr := `
	select user_id,userName,password,email,privilege,gender,avatar
	from user
	where user_id =?
`
	if err := mysql.Db.Get(user, sqlStr, userID); err != nil {
		return nil, err
	}
	return user, nil
}

func GetNoticeByID(NoticeType int64) (string, error) {
	var res string
	sqlStr := `
	select text
	from notice
	where id=?
`
	if err := mysql.Db.Get(&res, sqlStr, NoticeType); err != nil {
		return "", err
	}
	return res, nil
}

func UpdateUserByID(userID int64, user *User) error {
	sqlStr := `
	update user
	set avatar =? ,email=?
 	where user_id=?
`
	if _, err := mysql.Db.Exec(sqlStr, user.Avatar, user.Email, user.ID); err != nil {
		return err
	}
	return nil
}
