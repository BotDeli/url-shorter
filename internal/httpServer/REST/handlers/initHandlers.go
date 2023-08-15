package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"url-shorter/internal/database/mongodb"
	"url-shorter/pkg/cache"
)

func InitHandlers(router *gin.Engine, logger *slog.Logger, storage mongodb.Storage, cache cache.Cache, address string) {
	router.GET("/", showMainPage)

	router.POST("/getShortUrl", HandlerGetUrl(storage, logger, cache, address))

	router.NoRoute(HandlerRedirectUrl(logger, storage))
}
