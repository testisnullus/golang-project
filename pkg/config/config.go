package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type YamlFile struct {
	ServerPort string `yaml:"serverPort"`
	Config     struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbName"`
		SSLMode  string `yaml:"sslmode"`
	}
	ApiKey string     `yaml:"apiKey"`
	JWT    *JWTConfig `yaml:"jwt"`
}

type JWTConfig struct {
	HmacSecret string        `yaml:"hmacSecret"`
	Lifetime   time.Duration `yaml:"lifetime"`
}

func GetConfig(path string) (*YamlFile, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &YamlFile{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func PostgresURL(yaml *YamlFile) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", yaml.Config.Username, yaml.Config.Password, yaml.Config.Host, yaml.Config.Port, yaml.Config.DbName, yaml.Config.SSLMode)
}
