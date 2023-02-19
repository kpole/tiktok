package db

import (
	favorite "offer_tiktok/biz/model/interact/favorite"
	//relation "offer_tiktok/biz/model/social/relation"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type Favorites struct {
	ID        int64          `json:"id"`
	UserId    int64          `json:"user_id"`
	VideoId   int64          `json:"video_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (f *Favorites) TableName() string {
	return constants.FavoritesTableName
}

func AddNewFavorite(favorite *Favorites) (bool, error) {
	err := DB.Create(favorite).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteFavorite(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).Delete(favorite).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func QueryFavoriteExist(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).Find(&favorite).Error

	if err != nil {
		return false, err
	}
	if favorite.ID == 0 {
		return false, nil
	}
	return true, nil
}

// 查询视频的点赞数量
func GetFavoriteCount(video_id int64) (int64, error) {
	var count int64
	err := DB.Model(&Favorites{}).Where("video_id = ?", video_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 获得 user_id 点赞的视频的 video_id
func GetFavoriteIdList(user_id int64) ([]int64, error) {
	var favorite_actions []Favorites
	err := DB.Where("user_id = ?", user_id).Find(&favorite_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favorite_actions {
		result = append(result, v.VideoId)
	}
	return result, nil
}

// 获得 video_id 所有点赞的人的 id
func GetFavoriterIdList(video_id int64) ([]int64, error) {
	var favorite_actions []Favorites
	err := DB.Where("video_id = ?", video_id).Find(&favorite_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favorite_actions {
		result = append(result, v.UserId)
	}
	return result, nil
}

//-----------------------Add By ----------------------------//

func CheckFavoriteRelationExist(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).Find(&favorite).Error
	if err != nil {
		return false, err
	}
	// find未找到符合条件的数据会返回空结构体，ID = 0
	if favorite.ID == 0 {
		err := errno.FavoriteRelationNotExistErr
		return false, err
	}
	return true, nil
}

// // 返回用户点赞的视频的Video信息
func GetFavoriteInfo(current_user_id int64, user_id int64) ([]favorite.Video, error) {
	var favoriteing []Favorites
	err := DB.Where("user_id = ?", user_id).Find(&favoriteing).Error
	if err != nil {
		return nil, err
	}
	var result []favorite.Video

	// for _,favorite_id := range favoriteing {
	// 	video_id=favorite_id
	// 	follow_info, err := QueryUserById(follow_id.FollowerId)
	// 	if err != nil {
	// 		return result, err
	// 	}
	// 	FollowCount, err := GetFollowCount(follow_info.ID)
	// 	FolloweeCount, err := GetFolloweeCount(follow_info.ID)
	// 	IsFollow, err := CheckFollowRelationExist(&Follows{UserId: current_user_id, FollowerId: follow_id.FollowerId})
	// 	resp := &favorite.Video{
	// 		Id:            favorite_info.,
	// 		Name:          follow_info.UserName,
	// 		FollowCount:   FollowCount,
	// 		FollowerCount: FolloweeCount,
	// 		IsFollow:      IsFollow,
	// 	}
	// 	result = append(result, *resp)
	// }
	return result, nil
}

//------------------------------------------end---------------------------------//
