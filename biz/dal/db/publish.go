package db

import "time"

type Video struct {
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func CreateVideo(video *Video) (Video_id int64, err error) {
	err = DB.Create(video).Error
	if err != nil {
		return 0, err
	}
	return video.ID, err
}
