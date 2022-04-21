package handler_model

type ContentActionRequest struct {
	ContentId string `json:"content_id,omitempty"`
	Action    int64  `json:"action,omitempty"`
}

type ContentActionResponse struct {
}
