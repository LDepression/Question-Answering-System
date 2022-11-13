package logic

import (
	"wenba/internal/dao/mysql/sqlx"
	"wenba/internal/global"
	"wenba/internal/model/request"
	"wenba/internal/pkg/app/errcode"
)

//CreateCommentForQuestion 发表评论
func CreateCommentForQuestion(p request.Comment) errcode.Err {
	if p.AuthorID == 0 {
		return errcode.ErrParamsNotValid
	}
	if err := sqlx.CreateCommentForQuestion(p); err != nil {
		global.Logger.Error("sqlx.CreateCommentForQuestion(p) failed")
		return errcode.ErrServer.WithDetails("sqlx.CreateCommentForQuestion(p)", err.Error())
	}
	return nil
}

//PostReply 发表评论
func PostReply(p request.Comment) errcode.Err {
	if p.AuthorID == 0 || p.ParentID == 0 {
		return errcode.ErrParamsNotValid
	}

	if err := sqlx.PostReply(p); err != nil {
		global.Logger.Error("sqlx.PostReply(p) failed")
		return errcode.ErrServer.WithDetails("sqlx.PostReply(p)", err.Error())
	}
	return nil
}

func DeleteComment(p request.Comment) errcode.Err {
	//现在判断level,如果是1的话,就要将1子目录的评论全部要删除,如果是2的话,可以直接删除
	/*
		if p.Level == 2 {
			err := sqlx.DeleteComment2(p.CommentID)
			if err != nil {
				return errcode.ErrServer.WithDetails(err.Error())
			}
		} else if p.Level == 1 {

		}
	*/
	return nil
}
