package db

import (
	"offer_tiktok/pkg/constants"
)

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func QueryUser(userName string) (*User, error) {
	var user User
	if err := DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
