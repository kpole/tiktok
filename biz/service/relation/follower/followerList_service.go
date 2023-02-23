package service

import (
	"context"
	"log"
	user_service "offer_tiktok/biz/service/user"

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
		user_info, err := user_service.NewUserService(s.ctx, s.c).GetUserInfo(follower, current_user_id.(int64))
		if err != nil {
			log.Printf("func error: GetFollowerList -> GetUserInfo")
		}

		user := &relation.User{
			Id:              user_info.Id,
			Name:            user_info.Name,
			FollowCount:     user_info.FollowCount,
			FollowerCount:   user_info.FollowerCount,
			IsFollow:        user_info.IsFollow,
			Avatar:          user_info.Avatar,
			BackgroundImage: user_info.BackgroundImage,
			Signature:       user_info.Signature,
			TotalFavorited:  user_info.TotalFavorited,
			WorkCount:       user_info.WorkCount,
			FavoriteCount:   user_info.FavoriteCount,
		}
		followerList = append(followerList, user)
	}
	return followerList, nil
}
