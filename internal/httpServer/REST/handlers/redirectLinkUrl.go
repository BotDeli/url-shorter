package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
)

type storageUrl interface {
	GetUrl(string) (string, error)
}

func HandlerRedirectUrl(logger *slog.Logger, storage storageUrl) func(c *gin.Context) {
	return func(c *gin.Context) {
		shortUrl := c.Request.URL.String()

		url, err := getUrl(logger, storage, shortUrl)
		if err != nil {
			http.NotFound(c.Writer, c.Request)
			return
		}
		c.Redirect(http.StatusMovedPermanently, url)
	}
}

func getUrl(logger *slog.Logger, storage storageUrl, shortUrl string) (string, error) {
	url, err := storage.GetUrl(shortUrl)
	if err != nil {
		logger.Warn("HandlerRedirectUrl.getUrl", err)
	}

	return url, err
}
