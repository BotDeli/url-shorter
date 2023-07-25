package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
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

func HandlerGetUrl(storage storageShortUrl, logger *slog.Logger, address string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		if err = checkContentJSON(c); err != nil {
			return
		}

		var url struct {
			Url string `json:"url"`
		}

		if err = decodeJSON(c, logger, url); err != nil {
			return
		}

		if err = checkHomeAddress(c, url.Url, address); err != nil {
			return
		}

		if err = checkCorrectLengthUrl(c, url.Url); err != nil {
			return
		}

		shortUrl, err := getShortUrl(c, logger, storage, url.Url)

		if err != nil {
			return
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

func decodeJSON(c *gin.Context, logger *slog.Logger, output interface{}) (err error) {
	if err = c.ShouldBindJSON(&output); err != nil {
		logger.Warn("HandlerGetUrl.ShouldBindJSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	return err
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
	if len(url) <= 5 {
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
