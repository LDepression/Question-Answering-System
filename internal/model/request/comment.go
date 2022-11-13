package request

type Comment struct {
	CommentID      int64  `json:"comment_id" db:"comment_id"`
	Content        string `json:"content" binding:"required" db:"content"`
	AuthorID       int64  `json:"author_id" db:"author_id"`
	QuestionID     int64  `json:"question_id" binding:"required" db:"question_id"`
	LikeCount      int64  `json:"like_count"  db:"like_count"`
	CommentCount   int64  `json:"comment_count" db:"comment_count"`
	ParentID       int64  `json:"parent_id"  binding:"required" db:"parent_id"`
	ReplyAuthorID  int64  `json:"reply_author_id" db:"reply_author_id"`
	ReplyCommentID int64  `json:"reply_comment_id" db:"reply_comment_id"`
	Level          int64  `json:"level" db:"level"`
}
