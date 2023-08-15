package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

const (
	errPathIsNotSet   = "config file path is not set"
	errFileIsNotExist = "config file is not exist"
	errReadConfigFile = "error read config file: %s"
)

func MustGetConfig() (cfg Config) {
	path := os.Getenv("configPath")
	checkSetPath(path)
	checkFileIsExist(path)
	readConfig(path, &cfg)
	return
}

func checkSetPath(path string) {
	if path == "" {
		log.Fatal(errPathIsNotSet)
	}
}

func checkFileIsExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(errFileIsNotExist)
	}
}

func readConfig(path string, cfg *Config) {
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf(errReadConfigFile, err)
	}
}
