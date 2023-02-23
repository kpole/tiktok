package service

import (
	"context"
	"offer_tiktok/biz/dal/db"
	relation "offer_tiktok/biz/model/social/relation"
	user_service "offer_tiktok/biz/service/user"
	"offer_tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	FOLLOW   = 1
	UNFOLLOW = 2
)

type RelationService struct {
	ctx context.Context
	c   *app.RequestContext
}

//  new RelationService
func NewRelationService(ctx context.Context, c *app.RequestContext) *RelationService {
	return &RelationService{ctx: ctx, c: c}
}

// follow action, include follow and unfollow
// request parameters:
// string token = 1;       // 用户鉴权token
// int64 to_user_id = 2;   // 对方用户id
// int32 action_type = 3;  // 1-关注，2-取消关注
func (r *RelationService) FollowAction(req *relation.DouyinRelationActionRequest) (flag bool, err error) {
	// 颁发和验证token的行为均交给jwt处理，当发送到handler层时，默认已通过验证
	// 只需要检查参数ToUserIdD的合法性
	_, err = db.CheckUserExistById(req.ToUserId)
	if err != nil {
		return false, err
	}
	if req.ActionType != FOLLOW && req.ActionType != UNFOLLOW {
		return false, errno.ParamErr
	}
	// 获取current_user_id
	current_user_id, _ := r.c.Get("current_user_id")
	// 不准自己关注自己
	if req.ToUserId == current_user_id.(int64) {
		return false, errno.ParamErr
	}
	new_follow_relation := &db.Follows{
		UserId:     current_user_id.(int64),
		FollowerId: req.ToUserId,
	}
	// 请求参数校验完毕，检查follow表中是否已经存在这两者的关系
	follow_exist, _ := db.CheckFollowRelationExist(new_follow_relation)
	if req.ActionType == FOLLOW {
		if follow_exist {
			return false, errno.FollowRelationAlreadyExistErr
		}
		flag, err = db.AddNewFollow(new_follow_relation)
		//增加redis缓存功能
	} else {
		if !follow_exist {
			return false, errno.FollowRelationNotExistErr
		}
		flag, err = db.DeleteFollow(new_follow_relation)
		// 增加redis缓存功能
	}
	return flag, err
}

// 获取登录用户关注的所有用户列表，需要注意的是这里的token是客户端当前用户，而user_id可以是任意用户
// request parameters:
// string token;       // 用户鉴权token
// int64  user_id;     // 用户id
func (r *RelationService) GetFollowList(req *relation.DouyinRelationFollowListRequest) (followerlist []relation.User, err error) {
	_, err = db.CheckUserExistById(req.UserId)
	if err != nil {
		return nil, err
	}

	var followList []relation.User
	// 获取current_user_id
	current_user_id, exists := r.c.Get("current_user_id")
	if !exists {
		current_user_id = int64(0)
	}
	dbfollows, err := db.GetFollowIdList(req.UserId)
	if err != nil {
		return followList, err
	}

	for _, follow := range dbfollows {
		user_info, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(follow, current_user_id.(int64))
		if err != nil {
			continue
		}
		user := relation.User{
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
		followList = append(followList, user)
	}
	return followList, nil
}
