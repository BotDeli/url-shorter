package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type storageShortUrl interface {
	GetShortUrl(string) (string, error)
}

const (
	contentKey  = "Content-Type"
	contentJSON = "application/json"
)

func HandlerGetUrl(storage storageShortUrl, logger *slog.Logger, address string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get(contentKey) == contentJSON {
			var inputUrl struct {
				Url string `json:"url"`
			}

			if err := c.ShouldBindJSON(&inputUrl); err != nil {
				logger.Warn("HandlerGetUrl.ShouldBindJSON:", err)
				c.JSON(400, gin.H{"error": err})
				return
			}

			url := inputUrl.Url

			if isHomeAddress(url, address) {
				c.JSON(200, gin.H{"url": address})
				return
			}

			if len(url) > 5 {
				shortUrl, err := storage.GetShortUrl(url)
				if err == nil {
					c.JSON(200, gin.H{"url": address + shortUrl})
					return
				} else {
					logger.Warn("HandlerGetUrl.GetShortUrl:", err)
				}
			}
		}
		c.AbortWithStatus(400)
	}
}

func isHomeAddress(url, address string) bool {
	return url == address || url == address[8:]
}
