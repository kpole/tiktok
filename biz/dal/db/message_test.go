package db

import (
	"fmt"
	"testing"
	"time"
)

func TestAddNewMessage(t *testing.T) {
	Init()
	// one_min_before, _ := time.ParseDuration("-1m")
	message := &Messages{
		ToUserId:   1001,
		FromUserId: 1004,
		Content:    "test: 1004 -> 1001, this is a message",
		// CreatedAt:  time.Now().Add(one_min_before),
	}
	is_succ, err := AddNewMessage(message)

	if err != nil {
		fmt.Println("err 2")
	}
	if !is_succ {
		fmt.Println("failed 1")
	}
	time.Sleep(time.Second)
	message = &Messages{
		ToUserId:   1004,
		FromUserId: 1001,
		Content:    "test: 1001 -> 1004, this is latest message",
		// CreatedAt:  time.Now(),
	}
	is_succ, err = AddNewMessage(message)
	if err != nil {
		fmt.Println("err 2")
	}
	if !is_succ {
		fmt.Println("failed 2")
	}
}

func TestGetMessageByIdPair(t *testing.T) {
	Init()
	user_id1, user_id2 := 1004, 1001
	// 假设过来的是毫秒
	pre_msg_timestamp := int64(1676819765000)
	pre_msg_time := time.Unix(pre_msg_timestamp/1000, pre_msg_timestamp%1000*1000000)
	fmt.Println(pre_msg_time)

	messages, err := GetMessageByIdPair(int64(user_id1), int64(user_id2), pre_msg_time)
	if err != nil {
		fmt.Println("get error")
	}
	for _, message := range messages {
		fmt.Printf("%v\n", message)
	}
	fmt.Println("OK")
}
