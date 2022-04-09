package game

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/game_dal"
	"github.com/oasis/game_boat/biz/model/game_model"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	"github.com/oasis/game_boat/storage"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type GetGameThreeListHandler struct{}

func GetGameThreeList(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetGameThreeList]"

	page := cast.ToInt64(ctx.Query("page"))
	pageSize := cast.ToInt64(ctx.Query("page_size"))
	key := ctx.Query("option_list")

	logs.Info(fmt.Sprintf("%v, key = %v, page = %v, pageSize = %v", method, key, page, pageSize))

	threeList := storage.GetThreeList()

	resp := &handler_model.GetGameThreeListResponse{}

	g := GetGameThreeListHandler{}
	if key == "all" {
		g.getGames(resp, "new_game", page, pageSize, threeList)
		g.getGames(resp, "hot_game", page, pageSize, threeList)
		g.getGames(resp, "recent_update", page, pageSize, threeList)
	} else {
		g.getGames(resp, key, page, pageSize, threeList)
	}

	logs.Info(fmt.Sprintf("%v key = %v resp.len = %v %v %v", method, key, len(resp.NewGame), len(resp.RecentUpdate), len(resp.HotGame)))
	response.Success(ctx, resp)
}

func (g GetGameThreeListHandler) getGames(resp *handler_model.GetGameThreeListResponse, key string, page, pageSize int64, threeList map[string][]uint) error {
	ids := g.getIds(page, pageSize, key, threeList)
	gameDetailMap, err := game_dal.GetGamesDetail(ids)
	if err != nil {
		return err
	}

	if key == "new_game" {
		resp.NewGame = make([]*game_model.GameInfo, 0)
		for _, id := range ids {
			resp.NewGame = append(resp.NewGame, gameDetailMap[id])
		}
	} else if key == "recent_update" {
		resp.RecentUpdate = make([]*game_model.GameInfo, 0)
		for _, id := range ids {
			resp.RecentUpdate = append(resp.RecentUpdate, gameDetailMap[id])
		}
	} else {
		resp.HotGame = make([]*game_model.GameInfo, 0)
		for _, id := range ids {
			resp.HotGame = append(resp.HotGame, gameDetailMap[id])
		}
	}

	return nil
}

func (g GetGameThreeListHandler) getIds(page, pageSize int64, key string, threeList map[string][]uint) []uint {
	offset := (page - 1) * pageSize
	res := make([]uint, 0)
	for i := offset + 1; i <= page*pageSize && i < int64(len(threeList[key])); i++ {
		res = append(res, threeList[key][i])
	}
	return res
}
