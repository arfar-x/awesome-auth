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

// Jwt configuration.
type Jwt struct {
	SecretKey         string
	ExpirationSeconds int
}

var Config *AppConfig

// InitConfig Initialize application configurations by environment variables.
func InitConfig() *AppConfig {
	if err := godotenv.Load(".env.app"); err != nil {
		panic("Could not load environment variables.")
	}

	if Config == nil {
		Config = &AppConfig{
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
			Jwt: Jwt{
				SecretKey:         os.Getenv("JWT_SECRET"),
				ExpirationSeconds: getIntEnv("JWT_EXPIRATION_SECONDS"),
			},
		}
	}

	return Config
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
