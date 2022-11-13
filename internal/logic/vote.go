package logic

import (
	"github.com/gin-gonic/gin"
	"wenba/internal/dao/redis"
	"wenba/internal/global"
	mid "wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app/errcode"
)

func VoteForAnswer(c *gin.Context, userID int64, vote request.ReqVote) errcode.Err {
	if err := redis.VoteForAnswer(userID, vote); err != nil {
		global.Logger.Error("redis.VoteForAnswer failed"+err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}
