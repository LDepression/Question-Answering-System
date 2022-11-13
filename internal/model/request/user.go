package request

import (
	"wenba/internal/global"
	"wenba/internal/pkg/app/errcode"
)

//ReqRegister 处理注册的参数
type ReqRegister struct {
	Username   string `json:"username" binding:"required,gte=1"`
	Password   string `json:"password" binding:"required,gte=1"`
	Email      string `json:"email" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

func (r *ReqRegister) Judge() errcode.Err {
	switch {
	case len(r.Username) < global.Settings.Rule.MinUsernameLen || len(r.Username) > global.Settings.Rule.MaxUsernameLen:
		return errcode.ErrUsername
	case len(r.Password) < global.Settings.Rule.MinPasswordLen || len(r.Password) > global.Settings.Rule.MaxPasswordLen:
		return errcode.ErrPassword
	}
	return nil
}

type UploadAvatar struct {
	Avatar string `json:"avatar" form:"avatar"`
}
type ReqLogin struct {
	UserID       string `json:"userID"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	AccessToken  string
	RefreshToken string
}

func (r *ReqLogin) Judge() errcode.Err {
	switch {
	case len(r.Username) < global.Settings.Rule.MinUsernameLen || len(r.Username) > global.Settings.Rule.MaxUsernameLen:
		return errcode.ErrUsername
	case len(r.Password) < global.Settings.Rule.MinPasswordLen || len(r.Password) > global.Settings.Rule.MaxPasswordLen:
		return errcode.ErrPassword
	}
	return nil
}

type SendEmail struct {
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	OperationType int64  `json:"operation_type" binding:"required"`
	//1绑定邮箱 2解绑邮箱 3改密码
}

type ValidEmail struct {
}
