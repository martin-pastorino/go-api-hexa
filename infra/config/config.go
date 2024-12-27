package config


var redisHost string = "localhost"
var redisPort  int = 6379
var redisPassword string  = ""


type Config struct {
	RedisHost string
	RedisPort int
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