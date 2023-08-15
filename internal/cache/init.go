package cache

import (
	"log"
	"url-shorter/pkg/cache"
)

func MustInitCache(capacity int) cache.Cache {
	c, err := cache.CreateLRUCache(capacity)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
