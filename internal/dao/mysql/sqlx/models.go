package sqlx

import "fmt"

//Privilege 用户的权限
type Privilege string

const (
	PrivilegeBAN    Privilege = "BAN"
	PrivilegeValue1 Privilege = "管理员"
	PrivilegeValue2 Privilege = "用户"
)

func (e *Privilege) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Privilege(s)
	case string:
		*e = Privilege(s)
	default:
		return fmt.Errorf("unsupported scan type for Privilege: %T", src)
	}
	return nil
}

type Gender string

const (
	GenderValue0 Gender = "男"
	GenderValue1 Gender = "女"
	GenderValue2 Gender = "未知"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type User struct {
	ID        int64     `json:"id" db:"user_id"`
	UserName  string    `json:"userName" db:"userName"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Privilege Privilege `json:"privilege" db:"privilege"`
	Gender    Gender    `json:"gender" db:"gender"`
	Avatar    string    `json:"avatar" db:"avatar"`
}

type Question struct {
	ID              int64  `json:"question_id" db:"question_id"`
	AuthorID        int64  `json:"author_id" db:"author_id"`
	QuestionContent string `json:"question_content" db:"question_content"`
	CategoryID      int64  `json:"categoryID" db:"category_id"`
}

type Category struct {
	ID           int64  `json:"id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
}

type QuestionDetails struct {
	Question Question `json:"question"`
	Category Category `json:"category"`
}

type Answer struct {
	AnswerID     int64  `json:"answerID" db:"answer_id"`
	Content      string `json:"content" db:"answer_content"`
	QuestionID   int64  `json:"questionID" db:"question_id"`
	AuthorID     int64  `json:"authorID" db:"author_id"`
	CommentCount int64  `json:"commentCount" db:"comment_count"`
	Favourite    int64  `json:"favourite" db:"favourite"`
}

type APIAnswerDetails struct {
	*Answer `json:"answer"`
	*User   `json:"user"`
	Score   int64 `json:"score"`
}
