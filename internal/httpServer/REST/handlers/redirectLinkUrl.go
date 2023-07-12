package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
)

type storageUrl interface {
	GetUrl(string) (string, error)
}

func HandlerRedirectUrl(storage storageUrl, logger *slog.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		shortUrl := c.Request.URL.String()
		if len(shortUrl) == 13 {
			url, err := storage.GetUrl(shortUrl)
			if err == nil {
				c.Redirect(301, url)
				return
			} else {
				logger.Warn("HandlerRedirectUrl.GetUrl", err)
			}
		}
		notFound(c)
	}
}

func notFound(c *gin.Context) {
	http.NotFound(c.Writer, c.Request)
}
