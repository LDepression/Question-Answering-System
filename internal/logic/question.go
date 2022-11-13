package logic

import (
	"github.com/gin-gonic/gin"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app/errcode"
)

func QuestionSubmit(c *gin.Context, question request.ReqQuestion) errcode.Err {
	//将相关信息存到数据库中去
	err := sqlx.InsertQuestion(question)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}
func GetQuestionByID(c *gin.Context, questionID int64) (sqlx.Question, errcode.Err) {
	data, err := sqlx.GetQuestionByID(questionID)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return data, errcode.ErrServer
	}
	return data, nil
}

func GetAllQuestion(c *gin.Context, page int64) ([]sqlx.Question, errcode.Err) {
	data, err := sqlx.GetAllQuestion(int32(page))
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return data, nil
}

/*
func GetAllQuestionByDetails(c *gin.Context, page int32, questionID int64) errcode.Err {
	//先根据questionID查询出来问题的具体信息

	//这里先去获取问题的id
	data, err := sqlx.GetAllQuestionByDetails(page, questionID)
}
*/

//DeleteQuestionByCategoryID 通过categoryID去删除问题
func DeleteQuestionByCategoryID(c *gin.Context, categoryID int64) errcode.Err {
	if err := sqlx.DeleteQuestionByCategoryID(categoryID); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func GetQuestionListByCategoryID(c *gin.Context, categoryID int64) ([]sqlx.Question, errcode.Err) {
	data, err := sqlx.GetQuestionListByCategoryID(categoryID)
	if err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	return data, nil
}
