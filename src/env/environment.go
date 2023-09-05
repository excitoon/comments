package env

import "os"
import "strconv"

import _ "github.com/joho/godotenv/autoload"

var Host string
var Port int
var Endpoint string

var DatabaseHost string
var DatabasePort int
var DatabaseUser string
var DatabasePass string
var DatabaseName string
var DatabaseSchema string

func getEnvString(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intValue
}

func init() {
	Host = getEnvString("HOST", "0.0.0.0")
	Port = getEnvInt("PORT", 80)
	Endpoint = getEnvString("ENDPOINT", "/api")

	DatabaseHost = getEnvString("DB_HOST", "db")
	DatabasePort = getEnvInt("DB_PORT", 5432)
	DatabaseUser = getEnvString("DB_USER", "postgres")
	DatabasePass = getEnvString("DB_PASS", "postgres")
	DatabaseName = getEnvString("DB_NAME", "database")
	DatabaseSchema = getEnvString("DB_SCHEMA", "schema")
}
