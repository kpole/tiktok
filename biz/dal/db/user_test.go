package db

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Init()
	u := &User{
		UserName: "test",
		Password: "123456",
	}
	if err := CreateUser(u); err != nil {
		fmt.Printf("%v", false)
		return
	}
	fmt.Printf("%v", true)
}
