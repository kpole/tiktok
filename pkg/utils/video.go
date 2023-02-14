package utils

import "fmt"

func NewFileName(user_id int64, time int64) string {
	return fmt.Sprintf("%d.%d", user_id, time)
}
