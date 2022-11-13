package logic

import (
	"github.com/gin-gonic/gin"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/dao/redis"
	"wenba/internal/global"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/utils"
)

const DataLength = 2

func CreateAnswerForQuestion(c *gin.Context, answer request.ReqAnswer) errcode.Err {
	if err := redis.CreateAnswer(utils.StringToIDMust(utils.IDToSting(answer.ID))); err != nil {
		global.Logger.Error("redis.CreateAnswer"+err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	if err := sqlx.CreateAnswerForQuestion(answer); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func GetAnswerListByQuestionID(c *gin.Context, questionID int64) ([]sqlx.Answer, errcode.Err) {
	data, err := sqlx.GetAnswerListByQuestionID(questionID)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return data, errcode.ErrServer
	}
	answerIDs := make([]int64, 0, len(data))
	for _, v := range data {
		answerIDs = append(answerIDs, v.AnswerID)
	}
	return data, nil
}

func UpdateAnswerContent(c *gin.Context, answer request.ReqAnswer) errcode.Err {
	if err := sqlx.UpdateAnswerContent(answer); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func DeleteAnswerByAnswerID(c *gin.Context, answerID int64) errcode.Err {
	if err := sqlx.DeleteAnswerByAnswerID(answerID); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

//GetAnswerList2 获取回答列表的数据
func GetAnswerList2(list request.ReqAnswerList) ([]sqlx.APIAnswerDetails, errcode.Err) {
	ids, err := redis.GetAnswerIDsInOrder(list)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	//根据id去数据库中去查询答案的详情信息
	answerList, err := sqlx.GetAnswerListByIDs(ids)
	if err != nil {
		//注意这里如果err为空的话,调用这个withDetails的方法会爆空指针
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	//提前查询好answerIDs
	VoteNums, err := redis.GetVoteNum(ids)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	var answerDetails []sqlx.APIAnswerDetails
	//我们根据answerID得到回答的投票数
	for i, answerInfo := range answerList {
		AuthorID := answerInfo.AuthorID
		//根据作者ID去查询作者的相关信息
		user, _ := sqlx.GetUserByID(AuthorID)
		answerDetail := sqlx.APIAnswerDetails{
			Answer: &answerInfo,
			User:   user,
			Score:  VoteNums[i],
		}
		answerDetails = append(answerDetails, answerDetail)
	}
	return answerDetails, nil
}
