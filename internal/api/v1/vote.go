package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"wenba/internal/logic"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/utils"
)

//VoteForAnswer 为问题的回答投票
func VoteForAnswer(c *gin.Context) {
	rly := app.NewResponse(c)
	//先获取用户的投票参数
	//reqVote := new(request.ReqVote)
	var reqVote request.ReqVote
	if err := c.ShouldBindJSON(&reqVote); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			rly.Reply(errcode.ErrParamsNotValid)
			return
		}
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()), removeTopStruct(errs.Translate(trans)))
		return
	}
	//然后获取用户的id
	PayLoad, err := middleware.GetPayload(c)
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	userID := PayLoad.UserID
	//根据用户的id进行投票
	if err := logic.VoteForAnswer(c, utils.StringToIDMust(userID), reqVote); err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	rly.Reply(nil)
}
