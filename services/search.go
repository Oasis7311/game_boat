package services

import (
	"github.com/gin-gonic/gin"

	"github.com/oasis/game_boat/biz/dal/mysql/game_dal"
	"github.com/oasis/game_boat/biz/model/game_model"
)

type SearchService struct {
	ctx *gin.Context
}

func (s *SearchService) SearchGame(gameName string) ([]*game_model.GameInfo, error) {
	return game_dal.SearchGameByName(gameName)
}
