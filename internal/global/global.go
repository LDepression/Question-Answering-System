package global

import (
	ut "github.com/go-playground/universal-translator"
	"wenba/internal/model/config"
	"wenba/internal/pkg/app"
	email2 "wenba/internal/pkg/email"
	"wenba/internal/pkg/logger"
	"wenba/internal/pkg/snowflake"
)

var (
	RootDir   string     // 项目跟路径
	Settings  config.All //全局的配置
	Page      = new(app.Page)
	Logger    *logger.Log
	Snowflake *snowflake.Snowflake // 生成ID
	Email     = new(email2.Email)
	Trans     ut.Translator
)
