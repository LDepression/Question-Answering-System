package request

import (
	"wenba/internal/global"
	"wenba/internal/pkg/app/errcode"
)

type ReqQuestion struct {
	ID              int64  `json:"id" required:"binding"`
	AuthorID        int64  `json:"authorID"`
	QuestionContent string `json:"questionContent" required:"binding"`
	CategoryID      int64  `json:"categoryID" required:"binding"`
}
type GetQuestionPage struct {
	Page int32 `json:"page"`
}

func (r *GetQuestionPage) Judge() errcode.Err {
	var msg string
	switch {
	case r.Page > global.Settings.Page.MaxPageSize:
		msg = "页数超限"
	default:
		return nil
	}
	return errcode.ErrParamsNotValid.WithDetails(msg)
}
