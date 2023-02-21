package db

import (
	"fmt"
	"testing"
)

func TestAddNewComment(t *testing.T) {
	Init()
	comment := &Comment{
		UserId:      1000,
		VideoId:     115,
		CommentText: "video comment test",
	}
	err := AddNewComment(comment)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("insert succ")
}

func TestDeleteComment(t *testing.T) {
	Init()
	err := DeleteCommentById(1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("succ")
}
