package db

import (
	"offer_tiktok/pkg/constants"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(user *User) error {
	return DB.Create(user).Error
}
