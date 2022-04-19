package comment_model

type ListCommentDto struct {
	UserId    *uint
	GroupId   *uint
	RootId    *uint
	SortField string
	Desc      bool
	Status    int
	Offset    int
	Limit     int
}

func NewListCommentDto() *ListCommentDto {
	return &ListCommentDto{
		Status:    0,
		SortField: "create_time",
		Desc:      true,
	}
}
