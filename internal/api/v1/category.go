package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wenba/internal/logic"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app"
	"wenba/internal/pkg/app/errcode"
	"wenba/internal/pkg/utils"
)

func CreateCategory(c *gin.Context) {
	var reqCategory request.ReqCrCategory
	rly := app.NewResponse(c)
	if err := c.ShouldBindJSON(&reqCategory); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if err := logic.CreateCategory(c, reqCategory); err != nil {
		rly.Reply(errcode.ErrServer, "logic.CreateCategory() failed", err)
	}
	rly.Reply(nil)
}

func DeleteCategory(c *gin.Context) {
	rly := app.NewResponse(c)
	CategoryID := utils.StringToIDMust(c.Param("id"))
	fmt.Println(CategoryID)
	if err := logic.DeleteCategoryByID(c, CategoryID); err != nil {
		rly.Reply(err, "logic.DeleteCategoryByID failed", err)
	}
	rly.Reply(nil)
}
