package util

import "os"

func GetEnvVariable(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue
}
