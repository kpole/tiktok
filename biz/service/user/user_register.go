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
	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal/db"
)

type UserRegisterService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

// CreateUser create user info.
func (s *UserRegisterService) UserRegister(req *user.DouyinUserRegisterRequest) (string, int64, error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return "", 0, err
	}
	if *user != (db.User{}) {
		return "", 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return "", 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	token := req.Username + req.Password
	user_id, err := db.CreateUser(&db.User{
		UserName: req.Username,
		Password: passWord,
	})
	return token, user_id, err
}