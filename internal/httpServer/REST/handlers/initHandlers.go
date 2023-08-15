package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"url-shorter/internal/config"
	"url-shorter/internal/database/mongodb"
	"url-shorter/pkg/cache"
)

func InitHandlers(router *gin.Engine, cfg config.HTTPServerConfig, logger *slog.Logger, storage mongodb.Storage, cache cache.Cache) {
	router.GET("/", showMainPage)

	router.POST("/getShortUrl", HandlerGetUrl(storage, cfg, logger, cache))

	router.NoRoute(HandlerRedirectUrl(logger, storage))
}
