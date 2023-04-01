package db

import (
	"fmt"
	"testing"
)

func TestGetFriendList(t *testing.T) {
	Init()
	followerList, err := GetFollowerIdList(1001)
	if err != nil {
		fmt.Println("false")
		return
	}
	for _, followerId := range followerList {
		isFriend, err := QueryFollowExist(1001, followerId)
		if err != nil {
			fmt.Println("false")
			return
		}
		if isFriend {
			fmt.Println(followerId)
		}
	}
}
