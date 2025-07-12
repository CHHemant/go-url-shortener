package main

import (
	"sync"
)

var (
	db   = make(map[string]string)
	lock sync.RWMutex
)

func SaveURL(code, url string) {
	lock.Lock()
	defer lock.Unlock()
	db[code] = url
}

func GetURL(code string) string {
	lock.RLock()
	defer lock.RUnlock()
	return db[code]
}
