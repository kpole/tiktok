package service

import (
	"context"
	"log"
	user_service "offer_tiktok/biz/service/user"

	"github.com/cloudwego/hertz/pkg/app"

	"offer_tiktok/biz/dal/db"
	relation "offer_tiktok/biz/model/social/relation"
)

type FriendListService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFriendListService(ctx context.Context, c *app.RequestContext) *FriendListService {
	return &FriendListService{ctx: ctx, c: c}
}

// 相互关注的两个人互为好友
func (s *FriendListService) GetFriendList(req *relation.DouyinRelationFriendListRequest) ([]*relation.FriendUser, error) {
	user_id := req.UserId
	// token := req.Token
	var friendList []*relation.FriendUser

	// 首先获得 user_id 所有粉丝
	followerIdList, err := db.GetFollowerIdList(user_id)
	if err != nil {
		return friendList, err
	}

	for _, id := range followerIdList {
		// 查看 user_id 是否也关注了 他的粉丝
		isFriend, err := db.QueryFollowExist(&db.Follows{UserId: user_id, FollowerId: id})
		if err != nil {
			return friendList, err
		}
		if isFriend {
			user_info, err := user_service.NewUserService(s.ctx, s.c).GetUserInfo(id, user_id)
			if err != nil {
				log.Printf("func error")
			}
			friendList = append(friendList, &relation.FriendUser{
				Friend: &relation.User{
					Id:              user_info.Id,
					Name:            user_info.Name,
					FollowCount:     user_info.FollowCount,
					FollowerCount:   user_info.FollowerCount,
					IsFollow:        user_info.IsFollow,
					Avatar:          user_info.Avatar,
					BackgroundImage: user_info.BackgroundImage,
					Signature:       user_info.BackgroundImage,
					TotalFavorited:  user_info.TotalFavorited,
					WorkCount:       user_info.WorkCount,
					FavoriteCount:   user_info.FavoriteCount,
					// 因为相互关注的两个人互为好友，所以在已经是好友的情况下，必定关注了他
				},
				// 这两个字段的获得可能得有其它同学的聊天记录消息接口
				Message: "this is a message",
				MsgType: -1,
			})
		}
	}

	return friendList, nil
}
