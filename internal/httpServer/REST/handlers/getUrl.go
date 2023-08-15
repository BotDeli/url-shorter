package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"strings"
	"url-shorter/internal/config"
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
	ErrorContentType = errors.New("error content type")
	ErrorHomeAddress = errors.New("home address")
)

func HandlerGetUrl(storage storageShortUrl, cfg config.HTTPServerConfig, logger *slog.Logger, cache cache.Cache) gin.HandlerFunc {
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

		if !isValidUrl(url) {
			return
		}

		homeAddress := cfg.GetHomeAddress()
		if err = checkHomeAddress(c, url, homeAddress); err != nil {
			return
		}

		if !haveProtocol(url) {
			url = "http://" + url
		}

		if shortUrl, added = cache.Get(url); !added || shortUrl == "" {

			shortUrl, err = getShortUrl(c, logger, storage, url)

			if err != nil {
				return
			}
			_ = cache.Set(url, shortUrl)
		}
		c.JSON(http.StatusOK, gin.H{"url": "http://" + homeAddress + shortUrl})
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

func isValidUrl(url string) bool {
	return url != "http://" && url != "https://" && url != ""
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

func haveProtocol(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func getShortUrl(c *gin.Context, logger *slog.Logger, storage storageShortUrl, url string) (string, error) {
	shortUrl, err := storage.GetShortUrl(url)

	if err != nil {
		logger.Warn("HandlerGetUrl.GetShortUrl:", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	return shortUrl, err
}
