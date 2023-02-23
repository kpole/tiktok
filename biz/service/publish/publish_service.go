package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/feed"
	"offer_tiktok/biz/model/basic/publish"
	feed_service "offer_tiktok/biz/service/feed"
	"strconv"

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
	uploadinfo, err := minio.PutToBucket(s.ctx, constants.MinioVideoBucketName, req.Data)
	hlog.CtxInfof(s.ctx, "video upload size:"+strconv.FormatInt(uploadinfo.Size, 10))
	PlayURL := constants.MinioVideoBucketName + "/" + req.Data.Filename
	buf, err := ffmpeg.GetSnapshot(utils.URLconvert(s.ctx, s.c, PlayURL))
	uploadinfo, err = minio.PutToBucketByBuf(s.ctx, constants.MinioImgBucketName, filename+".png", buf)
	hlog.CtxInfof(s.ctx, "image upload size:"+strconv.FormatInt(uploadinfo.Size, 10))
	if err != nil {
		hlog.CtxInfof(s.ctx, "err:"+err.Error())
	}
	_, err = db.CreateVideo(&db.Video{
		AuthorID:    user_id,
		PlayURL:     PlayURL,
		CoverURL:    constants.MinioImgBucketName + "/" + filename + ".png",
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
				Id:              item.Author.Id,
				Name:            item.Author.Name,
				FollowCount:     item.Author.FollowCount,
				FollowerCount:   item.Author.FollowerCount,
				Avatar:          item.Author.Avatar,
				BackgroundImage: item.Author.BackgroundImage,
				Signature:       item.Author.Signature,
				TotalFavorited:  item.Author.TotalFavorited,
				WorkCount:       item.Author.WorkCount,
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
