package REST

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"url-shorter/internal/config"
	"url-shorter/internal/database/mongodb"
	"url-shorter/internal/httpServer/REST/handlers"
)

func StartServer(cfg config.HTTPServerConfig, logger *slog.Logger, storage mongodb.Storage) {
	router := gin.Default()
	configRouter(router, logger, storage, cfg.Address)
	startHandleServer(router, cfg)

}

func configRouter(router *gin.Engine, logger *slog.Logger, storage mongodb.Storage, address string) {
	loadFiles(router)
	handlers.InitHandlers(router, logger, storage, address)
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
}

func startHandleServer(router *gin.Engine, cfg config.HTTPServerConfig) {
	srv := http.Server{
		Addr:              cfg.Address,
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	log.Fatal(srv.ListenAndServe())
}
