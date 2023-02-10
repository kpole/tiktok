package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	_ "offer_tiktok/biz/dal"
	"offer_tiktok/biz/dal/db"
	user "offer_tiktok/biz/model/basic/user"
	"offer_tiktok/pkg/errno"
)

type UserLoginService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{ctx: ctx}
}

// CreateUser create user info.
func (s *UserLoginService) UserLogin(req *user.DouyinUserLoginRequest) (string, int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return "", 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	user_id, err := db.VerifyUser(req.Username, passWord)

	if err != nil {
		return "", 0, err
	}

	if user_id == 0 {
		return "", 0, errno.AuthorizationFailedErr
	}

	token := req.Username + req.Password
	return token, user_id, nil
}
