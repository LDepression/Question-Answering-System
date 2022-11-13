package request

type ReqVote struct {
	//UserID
	AnswerID  int64 `json:"answer_id" binding:"required"`
	Direction int   `json:"direction" binding:"oneof=1 0 -1"`
}
