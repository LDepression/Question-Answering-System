package logic

import (
	"github.com/gin-gonic/gin"
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/middleware"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app/errcode"
)

func CreateCategory(c *gin.Context, category request.ReqCrCategory) errcode.Err {
	if err := sqlx.CreateCategory(category); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return errcode.StatusOK
}

func DeleteCategoryByID(c *gin.Context, categoryID int64) errcode.Err {
	if err := sqlx.DeleteCategoryID(categoryID); err != nil {
		global.Logger.Error(err.Error(), middleware.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return errcode.StatusOK
}
