package env

import (
	"os"
	"strconv"
)

func GetEnvString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return value
}
