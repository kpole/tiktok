package service

import (
	"context"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/common"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/errno"
	"offer_tiktok/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"

	user "offer_tiktok/biz/model/basic/user"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// UserRegister register user return user id.
func (s *UserService) UserRegister(req *user.DouyinUserRegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}

	passWord, err := utils.Crypt(req.Password)
	user_id, err = db.CreateUser(&db.User{
		UserName:        req.Username,
		Password:        passWord,
		Avatar:          constants.TestAva,
		BackgroundImage: constants.TestBackground,
		Signature:       constants.TestSign,
	})
	return user_id, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.DouyinUserRequest) (*common.User, error) {
	query_user_id := req.UserId
	current_user_id, exists := s.c.Get("current_user_id")
	if !exists {
		current_user_id = 0
	}
	return s.GetUserInfo(query_user_id, current_user_id.(int64))
}

// GetUserInfo
//
//	@Description: 根据当前用户 user_id 查询 query_user_id 的信息
//	@receiver *UserService
//	@param query_user_id int64
//	@param user_id int64 当前登陆用户 id，可能为 0
//	@return *user.User
//	@return error
func (s *UserService) GetUserInfo(query_user_id, user_id int64) (*common.User, error) {
	u := &common.User{}

	dbUser, err := db.QueryUserById(query_user_id)
	if err != nil {
		return u, err
	}
	WorkCount, err := db.GetWorkCount(query_user_id)
	if err != nil {
		return u, err
	}
	FollowCount, err := db.GetFollowCount(query_user_id)
	if err != nil {
		return u, err
	}
	FollowerCount, err := db.GetFollowerCount(query_user_id)

	var IsFollow bool
	if user_id != 0 {
		IsFollow, err = db.QueryFollowExist(user_id, query_user_id)
		if err != nil {
			return u, err
		}
	} else {
		IsFollow = false
	}
	FavoriteCount, err := db.GetFavoriteCountByUserID(query_user_id)
	if err != nil {
		return u, err
	}
	TotalFavorited, err := db.QueryTotalFavoritedByAuthorID(query_user_id)
	if err != nil {
		return u, err
	}

	u = &common.User{
		Id:              query_user_id,
		Name:            dbUser.UserName,
		FollowCount:     FollowCount,
		FollowerCount:   FollowerCount,
		IsFollow:        IsFollow,
		Avatar:          utils.URLconvert(s.ctx, s.c, dbUser.Avatar),
		BackgroundImage: utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage),
		Signature:       dbUser.Signature,
		TotalFavorited:  TotalFavorited,
		WorkCount:       WorkCount,
		FavoriteCount:   FavoriteCount,
	}
	return u, nil
}
