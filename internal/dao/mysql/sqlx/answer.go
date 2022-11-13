package sqlx

import (
	"github.com/jmoiron/sqlx"
	"strings"
	"wenba/internal/dao/mysql"
	"wenba/internal/model/request"
)

func CreateAnswerForQuestion(answer request.ReqAnswer) (err error) {
	tx, err := mysql.Db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
	}()
	sqlStr := `
	insert into answer(answer_id,question_id,author_id,favourite,comment_count,answer_content)
	values (?,?,?,?,?,?)
`
	_, err = tx.Exec(sqlStr, answer.ID, answer.QuestionID, answer.AuthorID, answer.Like, answer.CommentCount, answer.Content)
	if err != nil {
		return
	}
	sqlStr1 := `
	update question
	set answer_count=answer_count+1
	where question_id=?
`
	_, err = tx.Exec(sqlStr1, answer.QuestionID)
	if err != nil {
		return
	}
	err = tx.Commit()
	return err
}

func GetAnswerListByQuestionID(questionID int64) ([]Answer, error) {
	var AnswerList []Answer
	sqlStr := `
	select answer_id,question_id,author_id,favourite,comment_count,answer_content
	from answer
	where question_id =?
`
	if err := mysql.Db.Select(&AnswerList, sqlStr, questionID); err != nil {
		return nil, err
	}
	return AnswerList, nil
}

func GetAnswerInfoByAnswerID(answerID int64) (Answer, error) {
	var answerInfo Answer
	sqlStr := `
	select answer_id,question_id,author_id,favourite,comment_count,answer_content
	from answer
	where answer_id=?
`
	err := mysql.Db.Get(&answerInfo, sqlStr, answerID)
	if err != nil {
		return answerInfo, err
	}
	return answerInfo, err
}

func UpdateAnswerContent(answer request.ReqAnswer) (err error) {
	sqlStr := `
	update answer
	set answer_content=?
	where answer_id=?
`
	_, err = mysql.Db.Exec(sqlStr, answer.Content, answer.ID)
	return err
}

func DeleteAnswerByAnswerID(answerID int64) (err error) {
	sqlStr := `
	delete from
	answer
	where answer_id=?
`
	_, err = mysql.Db.Exec(sqlStr, answerID)
	return err

}

//GetAnswerListByIDs 根据问题的id列表
func GetAnswerListByIDs(ids []string) ([]Answer, error) {
	var AnswerList []Answer
	sqlStr := `
	select answer_id,question_id,author_id,favourite,comment_count,answer_content
	from answer
	where answer_id IN (?)
	order by find_in_set(answer_id,?)
`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = mysql.Db.Rebind(query)
	err = mysql.Db.Select(&AnswerList, query, args...)
	return AnswerList, err
}
