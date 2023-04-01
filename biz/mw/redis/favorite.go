package redis

import (
	"strconv"
)

const (
	likeSuffix  = ":like"
	likedSuffix = ":liked"
)

type (
	Favorite struct{}
)

func (f Favorite) AddLike(user_id, video_id int64) {
	add(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

func (f Favorite) AddLiked(user_id, video_id int64) {
	add(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

func (f Favorite) DelLike(user_id, video_id int64) {
	del(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

func (f Favorite) DelLiked(user_id, video_id int64) {
	del(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

func (f Favorite) CheckLike(user_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) CheckLiked(video_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix)
}

func (f Favorite) ExistLike(user_id, video_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

func (f Favorite) ExistLiked(user_id, video_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

func (f Favorite) CountLike(user_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) CountLiked(video_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix)
}

func (f Favorite) GetLike(user_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) GetLiked(video_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix)
}
