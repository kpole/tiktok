package db

import (
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

func AddNewComment(comment *Comment) error {
	if ok, _ := CheckUserExistById(comment.UserId); !ok {
		return errno.UserIsNotExistErr
	}
	if ok, _ := CheckVideoExistById(comment.VideoId); !ok {
		return errno.VideoIsNotExistErr
	}
	err := DB.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCommentById(comment_id int64) error {
	if ok, _ := CheckCommentExist(comment_id); !ok {
		return errno.CommentIsNotExistErr
	}
	comment := &Comment{}
	err := DB.Where("id = ?", comment_id).Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckCommentExist(comment_id int64) (bool, error) {
	comment := &Comment{}
	err := DB.Where("id = ?", comment_id).Find(comment).Error
	if err != nil {
		return false, err
	}
	if comment.ID == 0 {
		return false, nil
	}
	return true, nil
}
