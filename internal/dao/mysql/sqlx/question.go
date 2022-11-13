package sqlx

import (
	"wenba/internal/dao/mysql"
	"wenba/internal/global"
	"wenba/internal/model/request"
)

func InsertQuestion(question request.ReqQuestion) error {
	sqlStr := `
	insert into question(question_id,category_id,author_id,question_content)
	values(?,?,?,?)
`
	if _, err := mysql.Db.Exec(sqlStr, question.ID, question.CategoryID, question.AuthorID, question.QuestionContent); err != nil {
		return err
	}
	return nil
}
func GetAllQuestion(page int32) ([]Question, error) {
	var data []Question
	sqlStr := `
	select question_id,category_id,author_id,question_content
	from question
	limit ?,?
`
	if err := mysql.Db.Select(&data, sqlStr, (page-1)*global.Page.DefaultPageSize, global.Page.DefaultPageSize); err != nil {
		return nil, err
	}
	return data, nil
}

func GetQuestionByID(questionID int64) (Question, error) {
	var question Question
	sqlStr := `
	select question_id,author_id,category_id,question_content
	from question
	where question_id=?
`
	if err := mysql.Db.Get(&question, sqlStr, questionID); err != nil {
		return question, err
	}
	return question, nil
}

func DeleteQuestionByCategoryID(categoryID int64) error {
	sqlStr := `
	delete
	from question
	where category_id=?
`
	_, err := mysql.Db.Exec(sqlStr, categoryID)
	if err != nil {
		return err
	}
	return nil
}

func GetQuestionListByCategoryID(categoryID int64) ([]Question, error) {
	var questionList []Question
	sqlStr := `
	select question_id,category_id,author_id,question_content
	from question
	where category_id=?
`
	if err := mysql.Db.Select(&questionList, sqlStr, categoryID); err != nil {
		return nil, err
	}
	return questionList, nil
}
