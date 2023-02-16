package service

import (
	"context"
	"log"
	"sync"
	"time"

	"offer_tiktok/biz/dal/db"
	feed "offer_tiktok/biz/model/basic/feed"
	user_service "offer_tiktok/biz/service/user"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type FeedService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFeedService(ctx context.Context, c *app.RequestContext) *FeedService {
	return &FeedService{ctx: ctx, c: c}
}

func (s *FeedService) Feed(req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	resp := &feed.DouyinFeedResponse{}
	var lastTime time.Time
	if req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(req.LatestTime, 0)
	}

	current_user_id, exists := s.c.Get("current_user_id")
	// 如果当前用户没有登陆，则将 current_user_id 赋值为 0
	if !exists {
		current_user_id = int64(0)
	}

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		return resp, err
	}

	videos := make([]*feed.Video, 0, constants.VideoFeedCount)
	err = s.copyVideos(&videos, &dbVideos, current_user_id.(int64))
	if err != nil {
		return resp, nil
	}
	resp.VideoList = videos
	if len(dbVideos) != 0 {
		resp.NextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}
	return resp, nil
}

func (s *FeedService) copyVideos(result *[]*feed.Video, data *[]*db.Video, userId int64) error {
	for _, item := range *data {
		video := s.createVideo(item, userId)
		*result = append(*result, video)
	}
	return nil
}

/**
 * @description: 将 db.Video 拼接成 feed.Video
 * @param {*db.Video} data
 * @param {int64} userId
 * @return {*}
 */
func (s *FeedService) createVideo(data *db.Video, userId int64) *feed.Video {
	video := &feed.Video{
		Id:       data.ID,
		PlayUrl:  utils.URLconvert(s.ctx, s.c, data.PlayURL),
		CoverUrl: utils.URLconvert(s.ctx, s.c, data.CoverURL),
		Title:    data.Title,
	}

	var wg sync.WaitGroup
	wg.Add(4)

	// 获取作者信息
	go func() {
		author, err := user_service.NewUserService(s.ctx, s.c).GetUserInfo(data.AuthorID, userId)
		if err != nil {
			log.Printf("func error")
		}
		video.Author = &feed.User{
			Id:            author.Id,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      author.IsFollow,
		}

		wg.Done()
	}()

	// 获取视频点赞数量
	go func() {
		// TODO: favorite_service
		video.FavoriteCount = 0
		wg.Done()
	}()

	go func() {
		// TODO
		video.CommentCount = 0
		wg.Done()
	}()

	go func() {
		// TODO
		video.IsFavorite = false
		wg.Done()
	}()

	wg.Wait()
	return video
}
