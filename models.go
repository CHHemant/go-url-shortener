package main

import (
	"crypto/sha1"
	"encoding/base64"
)

func GenerateCode(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(bs)[:7]
}
