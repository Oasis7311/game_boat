package recommend_model

type RecommendData struct {
	UserId uint
	TagId  uint
	Rel    uint
}

func (r RecommendData) TableName() string {
	return "recommend_data"
}
