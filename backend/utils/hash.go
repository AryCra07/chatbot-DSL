package utils

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
