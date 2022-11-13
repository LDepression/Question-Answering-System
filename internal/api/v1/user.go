package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wenba/internal/logic"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/jwt"
	"wenba/internal/pkg/utils"
)

func RegisterHandler(c *gin.Context) {
	rly := app.NewResponse(c)
	reqRegister := request.ReqRegister{}
	err := c.ShouldBindJSON(&reqRegister)
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	//这是将judge的err传给客户端
	if err := reqRegister.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	if err := logic.InsertUser(c, &reqRegister); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, nil)
	return
}

func LoginHandler(c *gin.Context) {
	rly := app.NewResponse(c)
	reqLogin := request.ReqLogin{}
	if err := c.ShouldBindJSON(&reqLogin); err != nil {
		rly.Reply(errcode.ErrServer)
		return
	}
	//对参数进行校验
	if err := reqLogin.Judge(); err != nil {
		rly.Reply(err)
		return
	}
	LoginInfo, err := logic.Login(c, reqLogin)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, gin.H{
		"username":     LoginInfo.Username,
		"accessToken":  LoginInfo.AccessToken,
		"refreshToken": LoginInfo.RefreshToken,
	})
}

func RefreshTokenHandle(c *gin.Context) {
	rly := app.NewResponse(c)
	rt := c.Query("refresh_token")
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
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken, //这里设置的进入网站的token
		"refresh_token": rToken, //这里是刷新网站的token
	})
}

func UploadAvatar(c *gin.Context) {
	rly := app.NewResponse(c)
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar request.UploadAvatar
	payLoad, err := middleware.GetPayload(c)
	if err != nil {
		rly.Reply(errcode.ErrUnauthorizedTokenGenerate)
		return
	}
	if err := c.ShouldBind(&uploadAvatar); err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	err = logic.UpdateAvatar(c, payLoad.UserName, file, fileSize)
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	rly.Reply(nil)
}

func SendingEmail(c *gin.Context) {
	var sendEmail request.SendEmail
	rly := app.NewResponse(c)
	payLoad, err := middleware.GetPayload(c)
	fmt.Println(payLoad)
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	userID := utils.StringToIDMust(payLoad.UserID)
	err1 := c.ShouldBindJSON(&sendEmail)
	if err1 != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err1.Error()))
		return
	}
	if err2 := logic.SendEmail(c, userID, sendEmail); err2 != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err2.Error()))
		return
	}
	rly.Reply(nil)
}

func ValidEmail(c *gin.Context) {
	rly := app.NewResponse(c)
	var validEmail request.ValidEmail
	if err := c.ShouldBind(&validEmail); err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}

	fmt.Println("c.GetHeader:", c.GetHeader("Proxy-Authorization"))
	if err := logic.ValidEmail(c, c.GetHeader("Proxy-Authorization")); err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	rly.Reply(nil)

}
