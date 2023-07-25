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

		if len(shortUrl) != 13 {
			notFound(c)
		}

		url, err := getUrl(c, logger, storage, shortUrl)
		if err != nil {
			return
		}

		c.Redirect(301, url)
	}
}

func notFound(c *gin.Context) {
	http.NotFound(c.Writer, c.Request)
}

func getUrl(c *gin.Context, logger *slog.Logger, storage storageUrl, shortUrl string) (string, error) {
	url, err := storage.GetUrl(shortUrl)

	if err != nil {
		logger.Warn("HandlerRedirectUrl.getUrl", err)
		notFound(c)
	}

	return url, err
}
