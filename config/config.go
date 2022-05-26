package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB   *DBConfig
	Port string
}

type DBConfig struct {
	Dialect  string
	Endpoint string
	Username string
	Password string
	Name     string
	Charset  string
}

func getDBConfig() *DBConfig {
	// Configure DB from environment variables.
	dbEnvs := map[string]string{
		"DB_ENDPOINT": "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
	}

	for key := range dbEnvs {
		value, found := os.LookupEnv(key)
		if !found {
			log.Fatalf("ENV variable %s not set", key)
		}

		dbEnvs[key] = value
	}

	return &DBConfig{
		Dialect:  "mysql",
		Endpoint: dbEnvs["DB_ENDPOINT"],
		Username: dbEnvs["DB_USER"],
		Password: dbEnvs["DB_PASSWORD"],
		Name:     dbEnvs["DB_NAME"],
		Charset:  "utf8",
	}
}

func GetConfig() *Config {
	config := &Config{
		DB:   getDBConfig(),
		Port: "8080",
	}

	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	}

	return config
}
