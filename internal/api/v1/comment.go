package v1

import (
	"github.com/gin-gonic/gin"
	"html"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/logic"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/utils"
)

const CommentID = "comment_id"

//CreateCommentForQuestion 发表评论
func CreateCommentForQuestion(c *gin.Context) {
	rly := app.NewResponse(c)
	var p request.Comment
	if err := c.ShouldBindJSON(&p); err != nil {
		global.Logger.Error(err.Error())
		rly.Reply(errcode.ErrParamsNotValid)
		return
	}
	//为comment生成id
	p.CommentID = global.Snowflake.GetID()
	//防止xss泄露
	p.Content = html.EscapeString(p.Content)
	payLoad, err := middleware.GetPayload(c)
	if err != nil {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	p.AuthorID = utils.StringToIDMust(payLoad.UserID)
	if err := logic.CreateCommentForQuestion(p); err != nil {
		global.Logger.Error("logic.CreateCommentForQuestion(p) failed")
		rly.Reply(errcode.ErrServer)
		return
	}
	rly.Reply(nil)
}

//PostReply 回复评论
func PostReply(c *gin.Context) {
	rly := app.NewResponse(c)
	var p request.Comment
	if err := c.ShouldBindJSON(&p); err != nil {
		global.Logger.Error(err.Error())
		rly.Reply(errcode.ErrParamsNotValid)
		return
	}
	//为comment生成id
	p.CommentID = global.Snowflake.GetID()
	//防止xss泄露
	//根据replyCommentID,去查询replyCommentID 的作者是谁
	p.Content = html.EscapeString(p.Content)
	payLoad, err := middleware.GetPayload(c)
	if err != nil {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	p.AuthorID = utils.StringToIDMust(payLoad.UserID)
	if err := logic.PostReply(p); err != nil {
		global.Logger.Error("logic.CreateCommentForQuestion(p) failed")
		rly.Reply(errcode.ErrServer)
		return
	}
	rly.Reply(nil)
}

//DeleteReply 删除评论
func DeleteReply(c *gin.Context) {
	//先去获取要删除的评论
	rly := app.NewResponse(c)
	commentID := c.Query(CommentID)
	if commentID == "0" {
		rly.Reply(errcode.ErrParamsNotValid)
		return
	}
	//根据CommentID去查询comment的相关信息
	commentInfo, err := sqlx.GetCommentByID(utils.StringToIDMust(commentID))
	if err != nil {
		rly.Reply(errcode.ErrServer)
		return
	}
	if err := logic.DeleteComment(commentInfo); err != nil {
		global.Logger.Error("logic.DeleteComment(commentInfo) failed" + err.Error())
		rly.Reply(err)
	}
	rly.Reply(nil)
}
