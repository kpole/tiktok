package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/feed"
	"offer_tiktok/biz/model/basic/publish"
	feed_service "offer_tiktok/biz/service/feed"

	"offer_tiktok/biz/mw/ffmpeg"
	"offer_tiktok/biz/mw/minio"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/utils"
	"path"
	"time"
)

type PublishService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewPublishService NewCreatePublishService new CreatePublishService
func NewPublishService(ctx context.Context, c *app.RequestContext) *PublishService {
	return &PublishService{ctx: ctx, c: c}
}

func (s *PublishService) PublishAction(req *publish.DouyinPublishActionRequest) (err error) {
	v, _ := s.c.Get("current_user_id")
	title := s.c.PostForm("title")
	user_id := v.(int64)
	nowTime := time.Now()
	filename := utils.NewFileName(user_id, nowTime.Unix())
	req.Data.Filename = filename + path.Ext(req.Data.Filename)
	_, err = minio.PutToBucket(s.ctx, constants.MinioVideoBucketName, req.Data)
	videoURL, err := minio.GetObjURL(s.ctx, constants.MinioVideoBucketName, req.Data.Filename)
	buf, err := ffmpeg.GetSnapshot(videoURL.String())
	_, err = minio.PutToBucketByBuf(s.ctx, constants.MinioImgBucketName, filename+".png", buf)
	_, err = db.CreateVideo(&db.Video{
		AuthorID:    user_id,
		PlayURL:     constants.MinioVideoBucketName + "/" + req.Data.Filename,
		CoverURL:    constants.MinioVideoBucketName + "/" + filename + ".png",
		PublishTime: nowTime,
		Title:       title,
	})
	return err
}

func (s *PublishService) PublishList(req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp = &publish.DouyinPublishListResponse{}
	query_user_id := req.UserId
	current_user_id, exist := s.c.Get("current_user_id")
	if !exist {
		current_user_id = int64(0)
	}
	dbVideos, err := db.GetVideoByUserID(query_user_id)
	if err != nil {
		return resp, err
	}
	var videos []*feed.Video

	f := feed_service.NewFeedService(s.ctx, s.c)
	err = f.CopyVideos(&videos, &dbVideos, current_user_id.(int64))
	if err != nil {
		return resp, err
	}
	for _, item := range videos {
		video := publish.Video{
			Id: item.Id,
			Author: publish.User{
				Id:            item.Author.Id,
				Name:          item.Author.Name,
				FollowCount:   item.Author.FollowCount,
				FollowerCount: item.Author.FollowerCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		resp.VideoList = append(resp.VideoList, video)
	}
	return resp, nil
}
