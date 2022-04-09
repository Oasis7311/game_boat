package get

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/content_dal"
	"github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/content_model"
	"github.com/oasis/game_boat/biz/model/handler_model"
	"github.com/oasis/game_boat/biz/model/user_model"
	"github.com/oasis/game_boat/global"
	utils2 "github.com/oasis/game_boat/utils"
	"github.com/oasis/game_boat/utils/response"
)

type CommunityHandler struct {
	contentMap       map[int64]*content_model.Content
	contentDetailMap map[int64]*content_model.ContentResponse
	contentBufferMap map[int64]*content_model.ContentBuffer
	authorInfoMap    map[int64]*user_model.UserInfo
}

func ListCommunity(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[ListMainPage]"
	req := new(handler_model.GetCommunityRequest)
	s := CommunityHandler{
		contentMap:       make(map[int64]*content_model.Content),
		contentDetailMap: make(map[int64]*content_model.ContentResponse),
		contentBufferMap: make(map[int64]*content_model.ContentBuffer),
		authorInfoMap:    make(map[int64]*user_model.UserInfo),
	}

	err := ctx.Bind(req)
	if err != nil {
		logs.Error(fmt.Sprintf("%v get req fail, err = %v", method, err))
		response.ValidateFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v req = %v", method, utils2.JsonStrFormatIgnoreErr(req)))

	resp := new(handler_model.GetCommunityResponse)

	contentIds, err := content_dal.GetAllContentId()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v GetAllContentId success, lend = %v", method, len(contentIds)))

	oldSlice := make([]interface{}, 0)
	for _, id := range contentIds {
		oldSlice = append(oldSlice, id)
	}
	newSlice := utils2.FakeShuffleNumSlice(oldSlice, req.Timestamp)

	needContentIds := make([]int64, 0)
	for i := (req.Page-1)*req.PageSize + 1; i <= req.Page*req.PageSize; i++ {
		if int(i) >= len(contentIds) {
			break
		}
		needContentIds = append(needContentIds, cast.ToInt64(newSlice[i]))
	}

	s.contentMap, err = content_dal.GetContentRaw(needContentIds)
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	logs.Info(fmt.Sprintf("%v get raw content success", method))

	err = s.getContentDetail()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	for _, id := range needContentIds {
		resp.ContentList = append(resp.ContentList, s.contentDetailMap[id])
	}
	resp.Timestamp = req.Timestamp

	response.Success(ctx, resp)
	logs.Info(fmt.Sprintf("%v success, contentIdList = %v, page = %v, pageSize = %v", method, needContentIds, req.Page, req.PageSize))

}

func (c *CommunityHandler) getContentDetail() error {
	if err := c.getContentAuthorInfo(); err != nil {
		return err
	}

	c.getContentBuffer()

	for contentId, content := range c.contentMap {
		c.contentDetailMap[contentId] = &content_model.ContentResponse{
			ContentId:     cast.ToString(contentId),
			AuthorInfo:    c.authorInfoMap[contentId],
			Title:         content.Title,
			SubTitle:      content.SubTitle,
			CoverImageUrl: content.CoverImageUrl,
			ImageList:     c.contentBufferMap[contentId].ContentImage,
			Summary:       c.contentBufferMap[contentId].Summary,
			PublishTime:   c.contentBufferMap[contentId].PublishTime,
			LikeCount:     0, //todo LikeCount填充
		}
	}

	return nil
}

func (c *CommunityHandler) getContentBuffer() {
	for id, content := range c.contentMap {
		buffer := &content_model.ContentBuffer{}
		json.Unmarshal(content.Buffer, buffer)
		c.contentBufferMap[id] = buffer
	}
}

func (c *CommunityHandler) getContentAuthorInfo() error {
	var err error
	c.authorInfoMap, err = user_dal.GetUserInfoByContent(c.contentMap)
	if err != nil {
		return err
	}
	return nil
}
