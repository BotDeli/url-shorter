package app

import (
	"url-shorter/internal/config"
	"url-shorter/internal/database/mongodb"
	"url-shorter/internal/httpServer/REST"
	"url-shorter/internal/logger"
)

func StartApplication() {
	cfg := config.MustGetConfig()
	logg := logger.MustStartLogger(cfg.Logger)
	storage := mongodb.MustNewStorage(cfg.Mongodb)
	defer storage.Disconnect()
	logg.Error("Start REST server:", REST.StartServer(cfg.HTTPServer, logg, storage))
}
