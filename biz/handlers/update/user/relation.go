package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	user_dal2 "github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

func PeopleRelate(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[PeopleRelate]"
	req := new(handler_model.PeopleRelateRequest)
	//s := new(RelationHandler)

	err := ctx.BindJSON(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	req.CommonRequest.UserId = cast.ToUint(ctx.Value("id"))

	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	if req.Relation > 0 {
		err = user_dal2.CreateAUserRelation(req.UserId, req.BUserId, req.Relation)
	} else {
		err = user_dal2.DeleteAUserRelation(req.UserId, req.BUserId, req.Relation*-1)
	}
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	response.Success(ctx, new(handler_model.PeopleRelateResponse))
}
