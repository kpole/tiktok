package db

import (
	"fmt"
	"testing"
)

func TestAddNewFavorite(t *testing.T) {
	Init()
	_, err := AddNewFavorite(&Favorites{
		UserId:  1000,
		VideoId: 115,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("succ")
}
