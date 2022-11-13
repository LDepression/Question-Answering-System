package request

type ReqAnswer struct {
	ID           int64  `json:"id"`
	Content      string `json:"content" binding:"required"`
	AuthorID     int64  `json:"authorID"`
	QuestionID   int64  `json:"questionID" binding:"required"`
	CommentCount int64  `json:"commentCount"`
	Like         int64  `json:"like"`
}

type ReqAnswerList struct {
	Page  int64  `json:"page" binding:"required"`
	Size  int64  `json:"size" binding:"required"`
	Order string `json:"order" binding:"required"`
}
