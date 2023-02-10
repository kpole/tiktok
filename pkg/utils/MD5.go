package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(str string) (string, error) {
	h := md5.New()
	if _, err := io.WriteString(h, str); err != nil {
		return "", err
	}
	data := fmt.Sprintf("%x", h.Sum(nil))
	return data, nil
}
