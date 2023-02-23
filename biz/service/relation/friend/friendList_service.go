package service

import (
	"context"
	"log"
	user_service "offer_tiktok/biz/service/user"

	"github.com/cloudwego/hertz/pkg/app"

	"offer_tiktok/biz/dal/db"
	relation "offer_tiktok/biz/model/social/relation"
	"offer_tiktok/pkg/errno"
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
	current_user_id, _ := s.c.Get("current_user_id")

	// 只有本人才可以看本人的好友列表
	if current_user_id.(int64) != user_id {
		return nil, errno.FriendListNoPremissionErr
	}

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
				log.Printf("func error: GetFriendList -> GetUserInfo")
			}
			message, err := db.GetLatestMessageByIdPair(user_id, id)

			if err != nil {
				log.Printf("func error: GetFriendList -> GetLatestMessageByIdPair")
			}
			var msgType int64
			if message == nil { // 自定义2表示双方无聊天记录
				msgType = 2
				message = &db.Messages{}
			} else if user_id == message.FromUserId {
				msgType = 1
			} else {
				msgType = 0
			}
			friendList = append(friendList, &relation.FriendUser{
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
				Message:         message.Content,
				MsgType:         msgType,
			})
		}
	}

	return friendList, nil
}
