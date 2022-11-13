package middleware

import (
	"strings"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/jwt"
	"wenba/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

type PayLoad struct {
	UserID    string         `json:"UserID"`
	UserName  string         `json:"UserName"`
	Privilege sqlx.Privilege `json:"Privilege"`
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		rly := app.NewResponse(c)
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(global.Settings.Token.AuthorizationKey, &PayLoad{
			UserID:    utils.IDToSting(int64(mc.UserID)),
			UserName:  mc.Username,
			Privilege: mc.Privilege,
		})
		c.Next() // 后续的处理请求的函数可以用过c.Get("ContextUserIDKey")来获取当前请求的用户信息
	}
}

// GetPayload 获取payload(前提是必须鉴权过)
func GetPayload(ctx *gin.Context) (*PayLoad, errcode.Err) {
	payload, ok := ctx.Get(global.Settings.Token.AuthorizationKey)
	if !ok {
		return nil, errcode.ErrUnauthorizedAuthNotExist
	}
	return payload.(*PayLoad), nil
}

// AuthMustManager 管理员校验,前提是登陆了
func AuthMustManager() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		payload, err := GetPayload(ctx)
		if err != nil {
			rly.Reply(err)
			ctx.Abort()
			return
		}
		if payload.Privilege != sqlx.PrivilegeValue1 {
			rly.Reply(errcode.ErrInsufficientPermissions)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
