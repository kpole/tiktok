package db

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Init()
	u := &User{
		UserName: "8",
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
	user, err := QueryUser("4")
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

func TestVerifyUser(t *testing.T) {
	Init()
	user_id, err := VerifyUser("test", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user_id)
}

func TestVerifyUser2(t *testing.T) {
	Init()
	user_id, err := VerifyUser("test", "1234523426")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user_id)
}
