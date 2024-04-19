package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig Application configuration.
type AppConfig struct {
	Name string
	Host string
	Port int
	Mode string
	DB   DB
	Jwt  Jwt
}

// DB Database configuration.
type DB struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
	Charset  string
}

type Jwt struct {
}

// InitConfig Initialize application configurations by environment variables.
func InitConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		panic("Could not load environment variables.")
	}

	return &AppConfig{
		Name: os.Getenv("APP_NAME"),
		Host: os.Getenv("APP_HOST"),
		Port: getIntEnv("APP_PORT"),
		Mode: os.Getenv("APP_MODE"),
		DB: DB{
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     getIntEnv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Charset:  os.Getenv("DB_CHARSET"),
		},
		Jwt: Jwt{},
	}
}

func getIntEnv(key string, defaultValue ...int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}
