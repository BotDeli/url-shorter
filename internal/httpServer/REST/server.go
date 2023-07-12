package REST

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"url-shorter/internal/config"
	"url-shorter/internal/database/mongodb"
	"url-shorter/internal/httpServer/REST/handlers"
)

func StartServer(cfg config.HTTPServerConfig, logger *slog.Logger, storage mongodb.Storage) error {
	router := gin.Default()
	loadFiles(router)

	router.GET("/", showMainPage)
	router.POST("/getShortUrl", handlers.HandlerGetUrl(storage, logger, cfg.Address))

	router.NoRoute(handlers.HandlerRedirectUrl(storage, logger))

	srv := http.Server{
		Addr:              cfg.Address,
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	err := srv.ListenAndServe()
	return err
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
}

func showMainPage(c *gin.Context) {
	c.HTML(200, "main_page.html", nil)
}
