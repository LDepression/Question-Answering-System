package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	version "wenba/internal/api/v1"
	mid "wenba/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(mid.GinLogger(), mid.LogBody())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pang",
		})
	})
	v1.POST("/register", version.RegisterHandler)
	v1.POST("/login", version.LoginHandler)
	v1.GET("/refreshToken", version.RefreshTokenHandle)

	v2 := v1.Group("/user")
	v2.Use(mid.JWTAuthMiddleware())
	{
		v2.POST("/uploadAvatar", version.UploadAvatar)
		v2.POST("/sendingEmail", version.SendingEmail)
		v2.POST("/validEmail", version.ValidEmail)
	}
	//问答模块
	cg := v1.Group("/question")
	cg.Use(mid.JWTAuthMiddleware())
	{
		cg.POST("/Submit", version.QuestionSubmit)
		cg.GET("/getAll", version.GetAllQuestion)
		//cg.GET("/getQuestionByDetails/:page/:question_id", version.GetQuestionByDetails)
		cg.GET("/getByID/:question_id", version.GetQuestionByID)
		//获取问题的详情
		cg.GET("/getDetailsByID/:question_id", version.GetDetailByID)
		cg.GET("/getQuestionListByCategoryID/:categoryID", version.GetQuestionListByCategoryID)
	}

	//分类相关
	cg1 := v1.Group("/category")
	cg1.Use(mid.AuthMustManager())
	{
		cg1.POST("/create", version.CreateCategory)
		cg1.DELETE("/delete/:id", version.DeleteCategory)
	}

	//回答问题接口
	cg2 := v1.Group("/answer")
	cg2.Use(mid.JWTAuthMiddleware())
	{
		cg2.POST("/CreateAnswerForQuestion", version.CreateAnswerForQuestion)
		cg2.GET("/GetAnswerListByQuestionID", version.GetAnswerListByQuestionID)
		cg2.POST("/GetAnswerList2", version.GetAnswerList2)
		cg2.PUT("/UpdateAnswerContent", version.UpdateAnswerContent)
		cg2.DELETE("/deleteAnswerByAnswerID/:answerID", version.DeleteAnswerByAnswerID)
	}
	//点赞与踩相关接口
	cg3 := v1.Group("/vote")
	cg3.Use(mid.JWTAuthMiddleware())
	{
		cg3.POST("/voteForAnswer", version.VoteForAnswer)
		//cg3.GET("/GetvoteNum", version.GetvoteNum)
	}

	//评论模块
	cg4 := v1.Group("/comment")
	cg4.Use(mid.JWTAuthMiddleware())
	{
		cg4.POST("/CreateCommentForQuestion", version.CreateCommentForQuestion)
		cg4.POST("/PostReply", version.PostReply)
		cg4.DELETE("/DeleteReply", version.DeleteReply)
	}
	return r
}
