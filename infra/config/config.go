package config

import (
	"encoding/json"
	"fmt"
	"os"
)



type Config struct {
	RedisHost     string `json:"redis_host"`
	RedisPort     string `json:"redis_port"`
	RedisPassword string `json:"redis_password"`
}

func loadConfig() *Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	configFile := fmt.Sprintf("../infra/config/config.%s.json", env)
	file, err := os.Open(configFile)
	if err != nil {
		return nil
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil
	}

	return &config
}


func NewConfig() *Config {
	config := loadConfig()
	if config == nil {
		panic("failed to load config")
	}
	return config
}

func NewConfigProvider() *Config {
	return NewConfig()
}
