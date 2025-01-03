package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	executable, err := os.Executable()
	if err != nil {
		return nil
	}
	executableDir := filepath.Dir(executable)
	baseDir := filepath.Join(executableDir, "..") 
	configFile := filepath.Join(baseDir, "infra", "config", fmt.Sprintf("config.%s.json", env))

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("can't read config file", err)
		return nil
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("can't decode config JSON", err)
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
