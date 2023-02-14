package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/publish"
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
	nowTime := time.Now().Unix()
	//filename := utils.NewFileName(user_id, nowTime)
	//minio.PutToBucket()
	_, err = db.CreateVideo(&db.Video{
		AuthorID:    user_id,
		PlayURL:     "",
		CoverURL:    "",
		PublishTime: nowTime,
		Title:       title,
	})

	return err
}
