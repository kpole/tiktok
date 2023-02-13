package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"offer_tiktok/biz/dal/db"
	user "offer_tiktok/biz/model/basic/user"
	"offer_tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewCreateUserService new CreateUserService
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// CreateUser create user info.
func (s *UserService) UserRegister(req *user.DouyinUserRegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	user_id, err = db.CreateUser(&db.User{
		UserName: req.Username,
		Password: passWord,
	})
	return user_id, nil
}

func (s *UserService) UserInfo(req *user.DouyinUserRequest) (*user.User, error) {
	resp := &user.User{}
	query_user_id := req.UserId

	user_id, exists := s.c.Get("user_id")
	if !exists {
		return resp, errno.UserIsNotExistErr
	}

	u, err := db.QueryUserById(query_user_id)

	if err != nil {
		return resp, err
	}
	FollowCount, err := db.GetFollowCount(u.ID)
	FolloweeCount, err := db.GetFolloweeCount(u.ID)
	IsFollow, err := db.QueryFollowExist(&db.Follows{UserId: int64(user_id.(float64)), FollowerId: query_user_id})
	resp = &user.User{
		Id:            u.ID,
		Name:          u.UserName,
		FollowCount:   FollowCount,
		FollowerCount: FolloweeCount,
		IsFollow:      IsFollow,
	}
	return resp, nil
}
