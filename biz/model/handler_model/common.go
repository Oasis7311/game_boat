package handler_model

type CommonRequest struct {
	UserId uint `json:"id"`
}

type PageInfo struct {
	Page      int   `json:"page,omitempty"`
	PageSize  int   `json:"page_size,omitempty"`
	TotalPage int   `json:"total_page,omitempty"`
	TotalSize int64 `json:"total_size,omitempty"`
}
