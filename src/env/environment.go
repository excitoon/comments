package env

import (
	"os"
	"strconv"
	"time"
)

import _ "github.com/joho/godotenv/autoload"

var (
	Host     string
	Port     uint
	Endpoint string
)

var (
	DatabaseHost        string
	DatabasePort        uint
	DatabaseUser        string
	DatabasePass        string
	DatabaseName        string
	DatabaseSchema      string
	DatabaseMaxIdle     uint
	DatabaseMaxOpen     uint
	DatabaseIdleTimeout time.Duration
)

var (
	JwtSecretKey      []byte
	JwtRealm          string
	JwtAccessTimeout  time.Duration
	JwtRefreshTimeout time.Duration
)

func getString(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func getUint(key string, fallback uint) uint {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	uint64Value, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return fallback
	}

	return uint(uint64Value)
}

func init() {
	Host = getString("HOST", "0.0.0.0")
	Port = getUint("PORT", 80)
	Endpoint = getString("ENDPOINT", "/api")

	DatabaseHost = getString("DB_HOST", "db")
	DatabasePort = getUint("DB_PORT", 5432)
	DatabaseUser = getString("DB_USER", "postgres")
	DatabasePass = getString("DB_PASS", "postgres")
	DatabaseName = getString("DB_NAME", "database")
	DatabaseSchema = getString("DB_SCHEMA", "schema")
	DatabaseMaxIdle = getUint("DB_MAX_IDLE", 10)
	DatabaseMaxOpen = getUint("DB_MAX_OPEN", 100)
	DatabaseIdleTimeout = time.Duration(getUint("DB_IDLE_TIMEOUT", 3600)) * time.Second

	JwtSecretKey = []byte(getString("JWT_SECRET_KEY", "secret key"))
	JwtRealm = getString("JWT_REALM", "test zone")
	JwtAccessTimeout = time.Duration(getUint("JWT_ACCESS_TIMEOUT", 3600)) * time.Second
	JwtRefreshTimeout = time.Duration(getUint("JWT_REFRESH_TIMEOUT", 86400)) * time.Second
}
