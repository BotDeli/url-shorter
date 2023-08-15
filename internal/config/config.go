package config

type Config struct {
	Logger     LoggerConfig     `yaml:"logger" env-required:"true"`
	Mongodb    MongodbConfig    `yaml:"mongodb" env-required:"true"`
	HTTPServer HTTPServerConfig `yaml:"http-server" env-required:"true"`
	Cache      CacheConfig      `yaml:"cache" env-required:""`
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

type CacheConfig struct {
	Capacity int `yaml:"capacity" env-default:"100"`
}
