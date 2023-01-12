package config

import "github.com/caarlos0/env"

type Config struct {
	MongoURI  string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	JWTSecret string `env:"JWT_SECRET" envDefault:"secret"`
}

var Cfg *Config

func Load() error {
	Cfg = &Config{}
	return env.Parse(Cfg)
}
