package main

import (
	"crypto/sha1"
	"encoding/base64"
	"regexp"
)

func GenerateCode(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(bs)[:7]
}

func IsValidURL(url string) bool {
	re := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return re.MatchString(url)
}
