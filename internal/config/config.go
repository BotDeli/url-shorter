package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Logger     LoggerConfig     `yaml:"logger" env-required:"true"`
	Mongodb    MongodbConfig    `yaml:"mongodb" env-required:"true"`
	HTTPServer HTTPServerConfig `yaml:"http-server" env-required:"true"`
}

type LoggerConfig struct {
	Format string `yaml:"format" env-default:"text"`
	Level  string `yaml:"level" env-default:"info"`
}

type MongodbConfig struct {
	Uri        string `yaml:"uri" env-default:"mongodb://localhost:27017"`
	Database   string `yaml:"database" env-required:"true"`
	Collection string `yaml:"collection" env-required:"true"`
}

type HTTPServerConfig struct {
	Address           string        `yaml:"address" env-default:"localhost:8080"`
	ReadHeaderTimeout time.Duration `yaml:"read-header-timeout" env-default:"4"`
	IdleTimeout       time.Duration `yaml:"idle-timeout" env-default:"60"`
}

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
