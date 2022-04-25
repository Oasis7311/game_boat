package services

import (
	"encoding/json"
	"time"

	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/action_dal"
	"github.com/oasis/game_boat/biz/dal/mysql/comment_dal"
	"github.com/oasis/game_boat/biz/dal/mysql/user_dal"
	"github.com/oasis/game_boat/biz/model/content_model"
	"github.com/oasis/game_boat/biz/model/user_model"
)

type ContentDetailHandler struct {
	ContentMap       map[int64]*content_model.Content
	ContentDetailMap map[int64]*content_model.ContentDetail
	ContentBufferMap map[int64]*content_model.ContentBuffer
	AuthorInfoMap    map[int64]*user_model.UserInfo
}

func GetContentDetailHandler() *ContentDetailHandler {
	return &ContentDetailHandler{
		ContentMap:       make(map[int64]*content_model.Content),
		ContentDetailMap: make(map[int64]*content_model.ContentDetail),
		ContentBufferMap: make(map[int64]*content_model.ContentBuffer),
		AuthorInfoMap:    make(map[int64]*user_model.UserInfo),
	}
}

func (c *ContentDetailHandler) GetContentDetail() error {
	if err := c.GetContentAuthorInfo(); err != nil {
		return err
	}

	c.GetContentBuffer()
	for contentId, content := range c.ContentMap {
		c.ContentDetailMap[contentId] = &content_model.ContentDetail{
			ContentId:     cast.ToString(contentId),
			AuthorInfo:    c.AuthorInfoMap[contentId],
			Title:         content.Title,
			SubTitle:      content.SubTitle,
			CoverImageUrl: content.CoverImageUrl,
			ImageList:     c.ContentBufferMap[contentId].ContentImage,
			Text:          c.ContentBufferMap[contentId].Text,
			Summary:       c.ContentBufferMap[contentId].Summary,
			PublishTime:   c.ContentBufferMap[contentId].PublishTime,
			LikeCount:     cast.ToUint(c.getLikeCount(contentId)),
			CommentCount:  cast.ToUint(c.getCommentCount(contentId)),
		}
	}

	return nil
}

func (c *ContentDetailHandler) GetContentBuffer() {
	for id, content := range c.ContentMap {
		buffer := &content_model.ContentBuffer{}
		json.Unmarshal(content.Buffer, buffer)
		c.ContentBufferMap[id] = buffer
	}
}

func (c *ContentDetailHandler) GetContentAuthorInfo() error {
	var err error
	c.AuthorInfoMap, err = user_dal.GetUserInfoByContent(c.ContentMap)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContentDetailHandler) GetContentByContentDetail() {
	c.FillUpBuffer()
	for id, detail := range c.ContentDetailMap {
		c.ContentMap[id] = &content_model.Content{
			ContentId:     id,
			AuthorId:      detail.AuthorInfo.ID.ID,
			Title:         detail.Title,
			SubTitle:      detail.SubTitle,
			CoverImageUrl: detail.CoverImageUrl,
			Buffer:        make([]byte, 0),
		}
		if b, ok := c.ContentBufferMap[id]; ok {
			c.ContentMap[id].Buffer, _ = json.Marshal(b)
		}
	}
}

func (c *ContentDetailHandler) FillUpBuffer() {
	for id, detail := range c.ContentDetailMap {
		c.ContentBufferMap[id] = &content_model.ContentBuffer{
			ContentImage: detail.ImageList,
			Summary:      detail.Summary,
			PublishTime:  time.Now().Unix(),
			Text:         detail.Text,
		}
	}
}

func (c *ContentDetailHandler) getLikeCount(contentId int64) int64 {
	count, err := action_dal.GetLikeCount(contentId)
	if err != nil {
		return 0
	}
	return count
}

func (c *ContentDetailHandler) getCommentCount(contentId int64) int64 {
	return comment_dal.GetCommentCount(contentId)
}
