package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"url-shorter/internal/database/mongodb"
)

func InitHandlers(router *gin.Engine, logger *slog.Logger, storage mongodb.Storage, address string) {
	router.GET("/", showMainPage)

	router.POST("/getShortUrl", HandlerGetUrl(storage, logger, address))

	router.NoRoute(HandlerRedirectUrl(logger, storage))
}
