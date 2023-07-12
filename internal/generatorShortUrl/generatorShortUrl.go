package generatorShortUrl

import (
	"math/rand"
	"time"
)

const charsPattern = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortUrl() (url string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	limit := len(charsPattern)
	for i := 0; i < 12; i++ {
		n := random.Intn(limit)
		url += string(charsPattern[n])
	}
	return "/" + url
}
