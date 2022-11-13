package request

type ReqCrCategory struct {
	ID           int64  `json:"id" binding:"required"`
	CategoryName string `json:"categoryName" binding:"required"`
}
