package db

import (
	"fmt"
	"testing"
)

func TestAddNewFollow(t *testing.T) {
	Init()
	f := Follows{
		UserId:     1001,
		FollowerId: 1004,
	}
	ok, err := AddNewFollow(&f)
	if err != nil {
		fmt.Println("false")
		return
	}
	if ok {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failed!")
	}
	fmt.Println("true")
}

func TestDeleteFollow(t *testing.T) {
	Init()
	f := Follows{
		UserId:     1009,
		FollowerId: 1006,
	}
	ok, err := DeleteFollow(&f)
	if err != nil {
		fmt.Println("false")
		return
	}
	if ok {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failed!")
	}
	fmt.Println("true")
}

func TestGetFollowCount(t *testing.T) {
	Init()
	cnt, err := GetFollowCount(1009)
	if err != nil {
		fmt.Println("false")
		return
	}
	fmt.Println(cnt)
}

func TestGetFolloweeCount(t *testing.T) {
	Init()
	cnt, err := GetFolloweeCount(1001)
	if err != nil {
		fmt.Println("false")
		return
	}
	fmt.Println(cnt)
}

func TestGetFollowerList(t *testing.T) {
	Init()
	followerList, err := GetFollowerIdList(1001)
	if err != nil {
		fmt.Println("false")
		return
	}
	for _, followerId := range followerList {
		fmt.Println(followerId)
	}
}
