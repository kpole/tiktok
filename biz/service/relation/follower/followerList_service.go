package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"offer_tiktok/biz/dal/db"
	relation "offer_tiktok/biz/model/social/relation"
)

type FollowerListService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFollowerListService(ctx context.Context, c *app.RequestContext) *FollowerListService {
	return &FollowerListService{ctx: ctx, c: c}
}

func (s *FollowerListService) GetFollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*relation.User, error) {
	user_id := req.UserId
	// token := req.Token
	var followerList []*relation.User
	current_user_id, exists := s.c.Get("current_user_id")
	if !exists {
		current_user_id = int64(0)
	}

	dbfollowers, err := db.GetFollowerIdList(user_id)
	if err != nil {
		return followerList, err
	}

	for _, follower := range dbfollowers {
		user_info, err := db.QueryUserById(follower)
		if err != nil {
			return followerList, err
		}
		followCnt, err := db.GetFollowCount(user_info.ID)
		if err != nil {
			return followerList, err
		}
		followerCnt, err := db.GetFolloweeCount(user_info.ID)
		if err != nil {
			return followerList, err
		}
		_IsFollow, err := db.QueryFollowExist(&db.Follows{UserId: current_user_id.(int64), FollowerId: follower})
		if err != nil {
			return followerList, err
		}

		user := &relation.User{
			Id:            user_info.ID,
			Name:          user_info.UserName,
			FollowCount:   followCnt,
			FollowerCount: followerCnt,
			IsFollow:      _IsFollow,
		}
		followerList = append(followerList, user)
	}
	return followerList, nil
}
