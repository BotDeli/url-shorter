package REST

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"url-shorter/internal/config"
	"url-shorter/internal/database/mongodb"
	"url-shorter/internal/httpServer/REST/handlers"
	"url-shorter/pkg/cache"
)

func StartServer(cfg config.HTTPServerConfig, logger *slog.Logger, storage mongodb.Storage, cache cache.Cache) error {
	router := gin.Default()
	configRouter(router, cfg, logger, storage, cache)
	return startHandleServer(router, cfg)
}

func configRouter(router *gin.Engine, cfg config.HTTPServerConfig, logger *slog.Logger, storage mongodb.Storage, cache cache.Cache) {
	loadFiles(router)
	handlers.InitHandlers(router, cfg, logger, storage, cache)
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.Static("/images", "./images")
	router.LoadHTMLGlob("templates/*")
}

func startHandleServer(router *gin.Engine, cfg config.HTTPServerConfig) error {
	srv := http.Server{
		Addr:              cfg.GetHomeAddress(),
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	return srv.ListenAndServe()
}
