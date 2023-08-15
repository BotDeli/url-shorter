package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"url-shorter/pkg/cache"
)

type storageShortUrl interface {
	GetShortUrl(string) (string, error)
}

const (
	contentKey  = "Content-Type"
	contentJSON = "application/json"
)

var (
	ErrorContentType    = errors.New("error content type")
	ErrorHomeAddress    = errors.New("home address")
	ErrorShortLengthUrl = errors.New("error shorted length url")
)

func HandlerGetUrl(storage storageShortUrl, logger *slog.Logger, cache cache.Cache, address string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			url, shortUrl string
			added         bool
			err           error
		)

		if err = checkContentJSON(c); err != nil {
			return
		}

		if url, err = decodeUrl(c, logger); err != nil {
			return
		}

		if err = checkHomeAddress(c, url, address); err != nil {
			return
		}

		if err = checkCorrectLengthUrl(c, url); err != nil {
			return
		}

		if shortUrl, added = cache.Get(url); !added || shortUrl == "" {

			shortUrl, err = getShortUrl(c, logger, storage, url)

			if err != nil {
				return
			}
			fmt.Println(url)
			_ = cache.Set(url, shortUrl)
		}
		c.JSON(200, gin.H{"url": address + shortUrl})
	}
}

func checkContentJSON(c *gin.Context) error {
	if c.Request.Header.Get(contentKey) != contentJSON {
		c.AbortWithStatus(http.StatusBadRequest)
		return ErrorContentType
	}

	return nil
}

func decodeUrl(c *gin.Context, logger *slog.Logger) (string, error) {
	var (
		err error
		url = struct {
			Url string `json:"url"`
		}{}
	)
	if err = c.ShouldBindJSON(&url); err != nil {
		logger.Warn("HandlerGetUrl.ShouldBindJSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	return url.Url, err
}

func checkHomeAddress(c *gin.Context, url, address string) error {
	if isHomeAddress(url, address) {
		c.JSON(http.StatusOK, gin.H{"url": address})
		return ErrorHomeAddress
	}

	return nil
}

func isHomeAddress(url, address string) bool {
	return url == address || url == address[8:]
}

func checkCorrectLengthUrl(c *gin.Context, url string) error {
	if len(url) <= 10 {
		c.AbortWithStatus(http.StatusBadRequest)
		return ErrorShortLengthUrl
	}

	return nil
}

func getShortUrl(c *gin.Context, logger *slog.Logger, storage storageShortUrl, url string) (string, error) {
	shortUrl, err := storage.GetShortUrl(url)

	if err != nil {
		logger.Warn("HandlerGetUrl.GetShortUrl:", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	return shortUrl, err
}
