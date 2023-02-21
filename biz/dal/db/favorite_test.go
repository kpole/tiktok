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

func TestQueryTotalFavoritedByAuthorID(t *testing.T) {
	Init()
	author_id := 1000
	sum, err := QueryTotalFavoritedByAuthorID(int64(author_id))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}
