package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/logic"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/utils"
)

const QuestionDetailNum = 2

func QuestionSubmit(c *gin.Context) {
	rly := app.NewResponse(c)
	reqQuestion := request.ReqQuestion{}
	if err := c.ShouldBindJSON(&reqQuestion); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	payLoad, err := middleware.GetPayload(c)
	if err != nil {
		rly.Reply(err)
		return
	}
	user, _ := sqlx.GetUserByUsername(payLoad.UserName)
	if err != nil {
		rly.Reply(errcode.ErrServer)
		return
	}
	reqQuestion.AuthorID = user.ID
	if err := logic.QuestionSubmit(c, reqQuestion); err != nil {
		rly.Reply(errcode.ErrServer)
		return
	}
	rly.Reply(nil, nil)
}

func GetAllQuestion(c *gin.Context) {
	rly := app.NewResponse(c)
	page := c.Query("page")
	//var questionPage request.GetQuestionPage
	//if err := c.ShouldBindQuery(&questionPage); err != nil {
	//	rly.Reply(errcode.ErrServer)
	//}

	data, err := logic.GetAllQuestion(c, utils.StringToIDMust(page))
	if err != nil {
		rly.Reply(err)
	}
	rly.ReplyList(nil, data)
}

//GetQuestionByDetails 获取问题的详情

/*
func GetQuestionByDetails(c *gin.Context) {
	//应该是先得到分类的id 也就是categoryID
	rly := app.NewResponse(c)
	page := c.Param("page")
	questionID := c.Param("question_id")
	data, err := logic.GetAllQuestionByDetails(c, page, questionID)
	if err != nil {
		rly.Reply(errcode.ErrServer, "logic.GetAllQuestionByDetails failed", err)
	}
	rly.ReplyList(errcode.StatusOK, data)
}
*/

func GetQuestionByID(c *gin.Context) {
	rly := app.NewResponse(c)
	//先根据命令行参数获得问题的ID
	questionID := c.Param("question_id")
	data, err := logic.GetQuestionByID(c, utils.StringToIDMust(questionID))
	if err != nil {
		rly.Reply(err, "logic.GetQuestionByID(questionID)", err)
	}
	rly.Reply(nil, data)
}

//GetDetailByID 通过questionID获取问题详情
func GetDetailByID(c *gin.Context) {
	rly := app.NewResponse(c)
	questionID := c.Param("question_id")
	//根据用户ID查询问题
	data, err := logic.GetQuestionByID(c, utils.StringToIDMust(questionID))
	if err != nil {
		rly.Reply(err, "logic.GetQuestionByID(questionID)", err)
	}
	//现在根据data中的question_id查询
	category, err1 := sqlx.GetCategoryByID(data.CategoryID)
	if err1 != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
	}
	mp := make(map[string]interface{}, QuestionDetailNum)
	mp["question"] = data
	mp["category"] = category
	rly.Reply(nil, mp)
}

func GetQuestionListByCategoryID(c *gin.Context) {
	rly := app.NewResponse(c)
	categoryID := utils.StringToIDMust(c.Param("categoryID"))
	data, err := logic.GetQuestionListByCategoryID(c, categoryID)
	if err != nil {
		rly.Reply(err, errcode.ErrServer.WithDetails(err.Error()))
	}
	fmt.Println(data)
	rly.ReplyList(nil, data)
}
