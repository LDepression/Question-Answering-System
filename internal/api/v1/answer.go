package v1

import (
	"github.com/gin-gonic/gin"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/logic"
	"wenba/internal/middleware"
	"wenba/internal/model/common"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/utils"
)

const questionID = "questionID"
const answerID = "answerID"

//CreateAnswerForQuestion 为问题提交题解
func CreateAnswerForQuestion(c *gin.Context) {
	rly := app.NewResponse(c)
	//先去绑定参数
	var answer request.ReqAnswer
	if err := c.ShouldBindJSON(&answer); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	payload, err := middleware.GetPayload(c)
	if err != nil {
		rly.Reply(err.WithDetails(err.Error()))
	}
	answer.AuthorID = utils.StringToIDMust(payload.UserID)
	answer.CommentCount = 0
	answer.Like = 0
	answer.ID = global.Snowflake.GetID()

	err = logic.CreateAnswerForQuestion(c, answer)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func GetAnswerListByQuestionID(c *gin.Context) {
	rly := app.NewResponse(c)
	questionID := c.Query(questionID)
	data, err := logic.GetAnswerListByQuestionID(c, utils.StringToIDMust(questionID))
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.ReplyList(nil, data)
}

//func GetAnswerListByQuestionID2(c *gin.Context) {
//	rly := app.NewResponse(c)
//	questionID := c.Query(questionID)
//	data, err := logic.GetAnswerListByQuestionID(c, utils.StringToIDMust(questionID))
//	if err != nil {
//		rly.Reply(err)
//		return
//	}
//	rly.ReplyList(nil, data)
//}

//UpdateAnswerContent 用户更新回答
func UpdateAnswerContent(c *gin.Context) {
	var answer request.ReqAnswer
	rly := app.NewResponse(c)
	//先要获取用户的信息
	payLoad, err := middleware.GetPayload(c)
	answerID := utils.StringToIDMust(c.Query(answerID))
	if err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	//鉴定当前用户是不是写了这篇文章
	answerInfo, err1 := sqlx.GetAnswerInfoByAnswerID(answerID)
	if err != nil {
		rly.Reply(errcode.ErrServer, err1.Error())
		return
	}
	//不是该回答的作者的话,就不能修改
	if answerInfo.AuthorID != utils.StringToIDMust(payLoad.UserID) {
		rly.Reply(errcode.ErrUnCorrentAuthor, gin.H{
			"msg": "鉴权失败,你不是这个回答的作者",
		})
		return
	}
	//接下来就是获取修改的回答的内容了
	if err := c.ShouldBindJSON(&answer); err != nil {
		rly.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	answer.ID = answerID
	if err := logic.UpdateAnswerContent(c, answer); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func DeleteAnswerByAnswerID(c *gin.Context) {
	rly := app.NewResponse(c)
	answerID := utils.StringToIDMust(c.Param(answerID))
	//先要获取用户的信息
	payLoad, err := middleware.GetPayload(c)
	//鉴定当前用户是不是写了这篇文章
	answerInfo, err1 := sqlx.GetAnswerInfoByAnswerID(answerID)
	if err != nil {
		rly.Reply(errcode.ErrServer, err1.Error())
		return
	}
	//不是该回答的作者的话,就不能修改
	if answerInfo.AuthorID != utils.StringToIDMust(payLoad.UserID) {
		rly.Reply(errcode.ErrUnCorrentAuthor, gin.H{
			"msg": "鉴权失败,你不是这个回答的作者",
		})
		return
	}
	if err := logic.DeleteAnswerByAnswerID(c, answerID); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

//GetAnswerList 获取答案的列表
/*
1.获取参数
2.去redis查询答案的id列表
3.根据id去数据库中查询相关数据
*/

func GetAnswerList2(c *gin.Context) {
	rly := app.NewResponse(c)
	p := request.ReqAnswerList{
		Page:  0,
		Size:  int64(global.Page.DefaultPageSize),
		Order: common.OrderTime,
	}
	if err := c.ShouldBindJSON(&p); err != nil {
		global.Logger.Error("c.ShouldBindJSON(&p) failed"+err.Error(), middleware.ErrLogMsg(c)...)
		rly.Reply(errcode.ErrParamsNotValid)
		return
	}
	//获取数据
	data, err := logic.GetAnswerList2(p)
	if err != nil {
		rly.Reply(err, "logic.GetAnswerList2 failed")
		return
	}
	rly.Reply(nil, data)
}
