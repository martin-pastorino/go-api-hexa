package config

import "os"


var redisHost string = os.Getenv("REDIS_HOST")
var redisPort  string = os.Getenv("REDIS_PORT")
var redisPassword string  = ""


type Config struct {
	RedisHost string
	RedisPort string
	RedisPassword string
}

func NewConfig() *Config {
	return &Config{
		RedisHost: redisHost,
		RedisPort: redisPort,
		RedisPassword: redisPassword,
	}
}

func NewConfigProvider() *Config {
	return NewConfig()
}