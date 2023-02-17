package db

import (
	"fmt"
	"testing"
	"time"
)

func TestGetVideoByLastTime(t *testing.T) {
	Init()
	lastTime := time.Now()
	videos, err := GetVideosByLastTime(lastTime)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestGetVideoByUserID(t *testing.T) {
	Init()
	user_id := int64(1000)
	videos, err := GetVideoByUserID(user_id)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}
