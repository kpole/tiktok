package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/publish"

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

// NewCreatePublishService new CreatePublishService
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

