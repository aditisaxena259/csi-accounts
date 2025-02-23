package initializers

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL           string
	PORT             string
	JWT_SECRET       string
	REDIS_HOST       string
	REDIS_PORT       string
	REDIS_PASSWORD   string
}

var CONFIG Config

func LoadEnv() {
	if os.Getenv("USE_ENV_FILE") == "" || os.Getenv("USE_ENV_FILE") == "true" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	configType := reflect.TypeOf(CONFIG)
	for i := 0; i < configType.NumField(); i++ {
		fieldName := configType.Field(i).Name
		envValue := getEnv(fieldName)
		setConfigField(&CONFIG, fieldName, envValue)
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func setConfigField(config *Config, fieldName string, value string) {
	structValue := reflect.ValueOf(config).Elem()
	structFieldValue := structValue.FieldByName(fieldName)

	if structFieldValue.IsValid() && structFieldValue.CanSet() {
		structFieldValue.SetString(value)
	}
}
