package action_model

type Action struct {
	UserId  uint  `json:"user_id,omitempty"`
	GroupId int64 `json:"group_id,omitempty"`
}
