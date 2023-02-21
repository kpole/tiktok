/*
 * @Description:
 * @Author: kpole
 * @Date: 2023-02-20 21:30:04
 * @LastEditors: kpole
 */
package service

import (
	"context"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/interact/comment"
	user_service "offer_tiktok/biz/service/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type CommentService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewCommentService(ctx context.Context, c *app.RequestContext) *CommentService {
	return &CommentService{ctx: ctx, c: c}
}

func (c *CommentService) AddNewComment(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	current_user_id, _ := c.c.Get("current_user_id")
	video_id := req.VideoId
	action_type := req.ActionType
	comment_text := req.CommentText
	comment_id := req.CommentId
	comment := &comment.Comment{}
	// 发表评论
	if action_type == 1 {
		db_comment := &db.Comment{
			UserId:      current_user_id.(int64),
			VideoId:     video_id,
			CommentText: comment_text,
		}
		err := db.AddNewComment(db_comment)
		if err != nil {
			return comment, err
		}
		comment.Id = db_comment.ID
		comment.CreateDate = db_comment.CreatedAt.Format("2006-01-02 15:04:05")
		comment.Content = db_comment.CommentText
		comment.User, err = c.GetUserInfoById(current_user_id.(int64), current_user_id.(int64))
		if err != nil {
			return comment, err
		}
		return comment, nil
	} else {
		err := db.DeleteCommentById(comment_id)
		if err != nil {
			return comment, err
		}
		return comment, nil
	}
}

func (c *CommentService) GetUserInfoById(current_user_id int64, user_id int64) (*comment.User, error) {
	u, err := user_service.NewUserService(c.ctx, c.c).GetUserInfo(user_id, current_user_id)
	var comment_user *comment.User
	if err != nil {
		return comment_user, err
	}
	comment_user = &comment.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
	return comment_user, nil
}
