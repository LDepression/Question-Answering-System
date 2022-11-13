package sqlx

import (
	"fmt"
	"wenba/internal/dao/mysql"
	"wenba/internal/model/request"
)

func CreateCommentForQuestion(p request.Comment) error {
	tx, err := mysql.Db.Beginx()
	if err != nil {
		return err
	}
	sqlStr := `
	insert into comment(
	comment_id,content,author_id,
	like_count,comment_count
					)
	values(
		?,?,?,?,?
		)
`
	_, err = tx.Exec(sqlStr, p.CommentID, p.Content, p.AuthorID, p.LikeCount, p.CommentCount)
	if err != nil {
		tx.Rollback()
		return err
	}
	sqlStr = `
	insert into comment_rel(
	comment_id,parent_id,level,
	question_id,reply_author_id
			)
	values(
		?,?,?,?,?
		)
`
	_, err = tx.Exec(sqlStr, p.CommentID, p.ParentID, 1, p.QuestionID, p.ReplyAuthorID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func PostReply(p request.Comment) error {
	tx, err := mysql.Db.Beginx()
	if err != nil {
		return err
	}
	//根据replyID查询对应的作者
	var ReplyAuthorID int64

	sqlStr := `
	select author_id from comment where comment_id=?
`
	if err = tx.Get(&ReplyAuthorID, sqlStr, p.ReplyCommentID); err != nil {
		return err
	}
	if ReplyAuthorID == 0 {
		err = fmt.Errorf("invalid reply authorID")
		return err
	}
	p.ReplyAuthorID = ReplyAuthorID
	sqlStr = `
	insert into comment(
	comment_id,content,author_id,
	like_count,comment_count
					)
	values(
		?,?,?,?,?
		)
`
	_, err = tx.Exec(sqlStr, p.CommentID, p.Content, p.AuthorID, p.LikeCount, p.CommentCount)
	if err != nil {
		tx.Rollback()
		return err
	}
	sqlStr = `
	insert into comment_rel(
	comment_id,parent_id,level,
	question_id,reply_author_id,reply_comment_id
			)
	values(
		?,?,?,?,?,?
		)
`
	_, err = tx.Exec(sqlStr, p.CommentID, p.ParentID, 2, p.QuestionID, p.ReplyAuthorID, p.ReplyCommentID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetCommentByID(commentID int64) (request.Comment, error) {
	var comment request.Comment
	sqlStr := `
	select comment_id,parent_id,level,question_id,reply_author_id,reply_comment_id
	from comment_rel
	where comment_id=?
`
	if err := mysql.Db.Get(&comment, sqlStr, commentID); err != nil {
		return comment, err
	}
	sqlStr = `
		select content,author_id
		from comment
		where comment_id=?
	`
	if err := mysql.Db.Get(&comment, sqlStr, commentID); err != nil {
		return comment, err
	}
	return comment, nil
}

/*
func DeleteComment2(p request.Comment) {
	sqlStr := `
	delete from comment
	where comment_id=?
`

}
*/
