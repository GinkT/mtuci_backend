package config

import "os"

type Config struct {
	MongoUsername string
	MongoPassword string
}

func InitConfig() *Config {
	return &Config{
		MongoUsername: os.Getenv("MONGO_USER"),
		MongoPassword: os.Getenv("MONGO_PASSWORD"),
	}
}