package db

import (
	"offer_tiktok/pkg/constants"
	"time"

	"gorm.io/gorm"
)

// user_id 关注了 follower_id
type Follows struct {
	ID         int64          `json:"id"`
	UserId     int64          `json:"user_id"`
	FollowerId int64          `json:"follower_id"`
	CreatedAt  time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (f *Follows) TableName() string {
	return constants.FollowsTableName
}

func AddNewFollow(follow *Follows) (bool, error) {
	err := DB.Create(follow).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteFollow(follow *Follows) (bool, error) {
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Delete(follow).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func QueryFollowExist(follow *Follows) (bool, error) {
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Find(&follow).Error

	if err != nil {
		return false, err
	}
	if follow.ID == 0 {
		return false, nil
	}
	return true, nil
}

// 查询用户的关注数量
func GetFollowCount(user_id int64) (int64, error) {
	var count int64
	err := DB.Model(&Follows{}).Where("user_id = ?", user_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询用户的粉丝数量
// 前提是 follower_id 存在
func GetFolloweeCount(follower_id int64) (int64, error) {
	var count int64
	err := DB.Model(&Follows{}).Where("follower_id = ?", follower_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 获得 user_id 关注的人的 id
func GetFollowIdList(user_id int64) ([]int64, error) {
	var follow_actions []Follows
	err := DB.Where("user_id = ?", user_id).Find(&follow_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range follow_actions {
		result = append(result, v.FollowerId)
	}
	return result, nil
}

// 获得 follower_id 所有粉丝的 id
func GetFollowerIdList(follower_id int64) ([]int64, error) {
	var follow_actions []Follows
	err := DB.Where("follower_id = ?", follower_id).Find(&follow_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range follow_actions {
		result = append(result, v.UserId)
	}
	return result, nil
}
