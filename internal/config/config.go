package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     uint
	UserDB     string
	PasswordDB string
	DBName     string
	ServerPort uint
	DBSSLMode  string
}

type EnvType interface {
	string | uint
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	DBHost, err := GetEnv[string]("DB_HOST")
	if err != nil {
		return nil, err
	}

	DBPort, err := GetEnv[uint]("DB_PORT")
	if err != nil {
		return nil, err
	}

	UserDB, err := GetEnv[string]("USER_DB")
	if err != nil {
		return nil, err
	}

	PasswordDB, err := GetEnv[string]("PASSWORD_DB")
	if err != nil {
		return nil, err
	}

	DBName, err := GetEnv[string]("DB_NAME")
	if err != nil {
		return nil, err
	}

	ServerPort, err := GetEnv[uint]("SERVER_PORT")
	if err != nil {
		return nil, err
	}

	DBSSLMode, err := GetEnv[string]("DBSSL_MODE")
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     DBHost,
		DBPort:     DBPort,
		UserDB:     UserDB,
		PasswordDB: PasswordDB,
		DBName:     DBName,
		ServerPort: ServerPort,
		DBSSLMode:  DBSSLMode,
	}, nil
}

func GetEnv[T EnvType](key string) (T, error) {
	var zero T

	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return zero, fmt.Errorf("environment variable not found or empty: %s", key)
	}

	switch any(zero).(type) {
	case string:
		return any(value).(T), nil

	case uint:
		parsed, err := strconv.ParseUint(value, 10, 0)
		if err != nil {
			return zero, fmt.Errorf("invalid uint value for %s: %w", key, err)
		}

		return any(uint(parsed)).(T), nil

	default:
		return zero, fmt.Errorf("unsupported type for env var %s", key)
	}
}
