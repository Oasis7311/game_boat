package content

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

type GetContentDetailHandler struct {
	contentMap       map[int64]*content_model.Content
	contentDetailMap map[int64]*content_model.ContentResponse
	contentBufferMap map[int64]*content_model.ContentBuffer
	authorInfoMap    map[int64]*user_model.UserInfo
}

func GetContentDetail(ctx *gin.Context) {
	logs := utils2.NewLoggerWithXRId(ctx, global.App.Log)
	method := "[GetContentDetail]"
	req := new(handler_model.GetContentRequest)
	c := &GetContentDetailHandler{
		contentMap:       make(map[int64]*content_model.Content),
		contentDetailMap: make(map[int64]*content_model.ContentResponse),
		contentBufferMap: make(map[int64]*content_model.ContentBuffer),
		authorInfoMap:    make(map[int64]*user_model.UserInfo),
	}

	contentId, ok := ctx.GetQuery("content_id")
	if !ok {
		logs.Error(fmt.Sprintf("%v get query fail", method))
		response.BusinessFail(ctx, "get content_id fail")
		return
	}
	req.ContentId = cast.ToInt64(contentId)
	logs.Info(fmt.Sprintf("%v req = %v", utils2.JsonStrFormatIgnoreErr(req)))

	var err error
	c.contentMap, err = content_dal.GetContentRaw([]int64{req.ContentId})
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}

	resp := new(handler_model.GetContentResponse)

	err = c.getContentDetail()
	if err != nil {
		logs.Error(fmt.Sprintf("%v %+v", method, err))
		response.BusinessFail(ctx, err.Error())
		return
	}
	resp.Content = c.contentDetailMap[req.ContentId]

	response.Success(ctx, resp)
}

func (c *GetContentDetailHandler) getContentDetail() error {
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
			Text:          c.contentBufferMap[contentId].Text,
			PublishTime:   c.contentBufferMap[contentId].PublishTime,
			LikeCount:     0, //todo LikeCount填充
		}
	}

	return nil
}

func (c *GetContentDetailHandler) getContentBuffer() {
	for id, content := range c.contentMap {
		buffer := &content_model.ContentBuffer{}
		json.Unmarshal(content.Buffer, buffer)
		c.contentBufferMap[id] = buffer
	}
}

func (c *GetContentDetailHandler) getContentAuthorInfo() error {
	var err error
	c.authorInfoMap, err = user_dal.GetUserInfoByContent(c.contentMap)
	if err != nil {
		return err
	}
	return nil
}
