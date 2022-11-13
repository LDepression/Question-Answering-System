package common

// Pager 分页
type Pager struct {
	Page     int32 `json:"page" form:"page" default:"1"`                        // 第几页
	PageSize int32 `json:"page_size" form:"page_size" maximum:"10" default:"5"` // 每页大小
}
