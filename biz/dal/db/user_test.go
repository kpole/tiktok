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
	user_id, err := CreateUser(u)
	if err != nil {
		fmt.Printf("%v", false)
		return
	}
	fmt.Printf("%v", user_id)
}

func TestQueryUser(t *testing.T) {
	Init()
	user, err := QueryUser("test2")
	if err != nil {
		fmt.Println(false)
		return
	}

	fmt.Printf("%v", user)
}

func TestQueryUser2(t *testing.T) {
	Init()
	user, err := QueryUser("ttttttt")
	if err != nil {

	}
	if *user == (User{}) {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
