package services

import (
	"sort"

	"github.com/oasis/game_boat/biz/dal/mysql/recommend_dal"
)

type RecommendService struct {
	UserId       uint
	userRelTimes map[uint]*map[uint]int //userRelTimes[userId][rel] = n 用户rel的次数
	tagRelTimes  map[uint]*map[uint]int //tagRelTimes[tagId][rel] = n 标签被rel的次数
	userTagTimes map[uint]*map[uint]int //userTagTimes[userId][tagId] = n 用户与tagId发生rel的次数
}

// NewRecommendService 获取用户的推荐服务并初始化推荐数据
func NewRecommendService(userId uint) (*RecommendService, error) {
	s := &RecommendService{
		UserId:       userId,
		userRelTimes: make(map[uint]*map[uint]int),
		tagRelTimes:  make(map[uint]*map[uint]int),
		userTagTimes: make(map[uint]*map[uint]int),
	}
	err := s.initRecommendData()
	return s, err
}

//初始化推荐数据
func (r *RecommendService) initRecommendData() error {
	datas, err := recommend_dal.GetAllRecommendData(r.UserId)
	if err != nil {
		return err
	}
	for _, data := range datas {
		userId := data.UserId
		tagId := data.TagId
		rel := data.Rel
		r.addValueToMat(r.userRelTimes, userId, rel)
		r.addValueToMat(r.tagRelTimes, rel, tagId)
		r.addValueToMat(r.userTagTimes, userId, tagId)
	}
	return nil
}

func (r *RecommendService) GetRecommendGame() []uint {
	recommendList := make(map[uint]int)
	if _, ok := r.userRelTimes[r.UserId]; !ok {
		return make([]uint, 0)
	}
	UserTags := *r.userTagTimes[r.UserId]
	for rel, wut := range *r.userRelTimes[r.UserId] {
		for tag, wit := range *r.tagRelTimes[rel] {
			if _, ok := UserTags[tag]; !ok {
				recommendList[tag] += wut * wit
			} else {
				recommendList[tag] = wut * wit
			}
		}
	}

	type RecommendResult struct {
		TagId uint
		Score int
	}
	tRes := make([]RecommendResult, 0)
	for tagId, score := range recommendList {
		tRes = append(tRes, RecommendResult{TagId: tagId, Score: score})
	}
	sort.Slice(tRes, func(i, j int) bool {
		return tRes[i].Score > tRes[j].Score
	})
	res := make([]uint, 0)
	for _, result := range tRes {
		res = append(res, result.TagId)
	}
	return res
}

func (r *RecommendService) addValueToMat(mat map[uint]*map[uint]int, key, value uint) {
	if _, ok := mat[key]; !ok {
		mat[key] = new(map[uint]int)
	} else if _, ok := (*mat[key])[value]; !ok {
		(*mat[key])[value] = 1
	} else {
		(*mat[key])[value]++
	}

}
