package services

import (
	"crypto/md5"
	"fmt"
)

func (s *Services) passwordToHash(password string) string {
	h := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", h)
}
