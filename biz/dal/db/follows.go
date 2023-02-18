package db

import (
	relation "offer_tiktok/biz/model/social/relation"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/errno"
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

//-----------------------Add By BourneHUST----------------------------//

func CheckFollowRelationExist(follow *Follows) (bool, error) {
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Find(&follow).Error
	if err != nil {
		return false, err
	}
	// find未找到符合条件的数据会返回空结构体，ID = 0
	if follow.ID == 0 {
		err := errno.FollowRelationNotExistErr
		return false, err
	}
	return true, nil
}

// 返回用户关注者的User信息, 传入当前用户的id是为了判断isfollow字段
func GetFollowInfo(current_user_id int64, user_id int64) ([]relation.User, error) {
	var following []Follows
	err := DB.Where("user_id = ?", user_id).Find(&following).Error
	if err != nil {
		return nil, err
	}
	var result []relation.User
	//去user表获取关注用户的信息
	for _, follow_id := range following {
		follow_info, err := QueryUserById(follow_id.FollowerId)
		if err != nil {
			return result, err
		}
		FollowCount, err := GetFollowCount(follow_info.ID)
		FolloweeCount, err := GetFolloweeCount(follow_info.ID)
		IsFollow, err := CheckFollowRelationExist(&Follows{UserId: current_user_id, FollowerId: follow_id.FollowerId})
		resp := &relation.User{
			Id:            follow_info.ID,
			Name:          follow_info.UserName,
			FollowCount:   FollowCount,
			FollowerCount: FolloweeCount,
			IsFollow:      IsFollow,
		}
		result = append(result, *resp)
	}
	return result, nil
}

//------------------------------------------end---------------------------------//
