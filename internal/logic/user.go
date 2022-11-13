package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/jwt"
	"wenba/internal/pkg/password"
	"wenba/internal/pkg/utils"
)

//InsertUser 添加普通用户
func InsertUser(c *gin.Context, r *request.ReqRegister) errcode.Err {
	//先要去判断用户名是否已经存在
	err := sqlx.CheckUsername(r.Username)
	if !errors.Is(err, sql.ErrNoRows) {
		return errcode.ErrUsenameExist
	}
	//将密码进行加密
	hashPassword, err := password.HashPassword(r.Password)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	//生成用户id
	user := sqlx.User{}
	user.ID = global.Snowflake.GetID()
	user.UserName = r.Username
	user.Password = hashPassword
	user.Email = r.Email
	user.Gender = sqlx.GenderValue2
	user.Privilege = sqlx.PrivilegeValue2
	//将相关信息写进数据库中去
	if err := sqlx.InsertUser(user); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
	}
	return nil
}

//Login 判断用户的账号或密码是否正确
func Login(c *gin.Context, ParamLogin request.ReqLogin) (*request.ReqLogin, errcode.Err) {
	//hashPassword, err := password.HashPassword(ParamLogin.Password)
	//if err != nil {
	//	global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
	//	return nil, errcode.ErrServer
	//}
	user, err := sqlx.GetUserByUsername(ParamLogin.Username)
	//这里是将原来的密码加密后与hashPassword是否是一样的
	err = password.CheckPassword(ParamLogin.Password, user.Password)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	ParamLogin.Password = user.Password
	if err := sqlx.CheckUsernameAndPassword(ParamLogin); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	// 生成JWT
	//return jwt.GenToken(user.UserID,user.UserName)
	atoken, rtoken, err := jwt.GenToken(uint64(utils.StringToIDMust(ParamLogin.UserID)), ParamLogin.Username)
	if err != nil {
		return nil, errcode.ErrServer
	}
	ParamLogin.AccessToken = atoken
	ParamLogin.RefreshToken = rtoken
	return &ParamLogin, nil
}

//UpdateAvatar 头像更新
func UpdateAvatar(c *gin.Context, userName string, file multipart.File, fileSize int64) errcode.Err {
	user, err := sqlx.GetUserByUsername(userName)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	//保存图片到本地函数
	path, err := UploadAvatarToLocalStatic(file, user.ID, user.UserName)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	user.Avatar = path
	err = sqlx.UpdateUserByID(user.ID, user)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func SendEmail(c *gin.Context, userID int64, email request.SendEmail) errcode.Err {
	fmt.Println("email:", email)
	token, err := jwt.GenEmailToken(userID, email.OperationType, email.Email, email.Password)
	var address string
	var notice string
	notice, err1 := sqlx.GetNoticeByID(email.OperationType)
	if err1 != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err1.Error())
	}
	address = global.Settings.Email.ValidEmail + token //发送方
	mailStr := notice
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", global.Email.SmtpEmail)
	m.SetHeader("To", email.Email)
	m.SetHeader("subject", "wenba")
	m.SetBody("text/html", mailText) //设置body
	d := mail.NewDialer(global.Email.SmtpHost, 465, global.Email.SmtpEmail, global.Email.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return errcode.ErrServer.WithDetails("d.DialAndSend(m) failed", err.Error())
	}
	return nil
}

//ValidEmail 验证邮箱
func ValidEmail(c *gin.Context, token string) errcode.Err {
	var userID int64
	var email string
	//var password string
	var operationType int64
	if token == "" {
		return errcode.ErrParamsNotValid
	} else {
		claims, err := jwt.ParseEmailToken(token)
		if err != nil {
			return errcode.ErrUnauthorizedAuthNotExist
		} else if time.Now().Unix() > claims.ExpiresAt {
			return errcode.ErrUnauthorizedTokenTimeout
		} else {
			userID = claims.ID
			email = claims.Email
			//password = claims.Password
			operationType = claims.OperationType
		}
	}
	//获取该用户的信息
	user, err := sqlx.GetUserByID(userID)
	if err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	if operationType == 1 {
		//绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		user.Email = ""
	} else if operationType == 3 {
		//err = sqlx.SetPassword(password)
		//todo
		if err != nil {
			return errcode.ErrServer.WithDetails(err.Error())
		}
	}
	fmt.Println("user:", user)
	err = sqlx.UpdateUserByID(userID, user)
	if err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}
