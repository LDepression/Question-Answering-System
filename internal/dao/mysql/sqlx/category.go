package sqlx

import (
	"wenba/internal/dao/mysql"
	"wenba/internal/model/request"
)

func CreateCategory(category request.ReqCrCategory) error {
	sqlStr := `
	insert into category(category_id,category_name)
	values(?,?)
`
	if _, err := mysql.Db.Exec(sqlStr, category.ID, category.CategoryName); err != nil {
		return err
	}
	return nil
}

func GetCategoryByID(categoryID int64) (Category, error) {
	var category Category
	sqlStr := `
	select category_id,category_name
	from category
	where category_id=?
`
	if err := mysql.Db.Get(&category, sqlStr, categoryID); err != nil {
		return category, err
	}
	return category, nil
}

func DeleteCategoryID(categoryID int64) error {
	tx, err := mysql.Db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	sqlStr := `
	delete 
	from category
	where category_id=?
`
	_, err = tx.Exec(sqlStr, categoryID)
	if err != nil {
		tx.Rollback()
		return err
	}
	sqlStr1 := `
	delete
	from question
	where category_id=?
`
	if _, err := tx.Exec(sqlStr1, categoryID); err != nil {
		return err
	}
	tx.Commit()
	return nil
}
