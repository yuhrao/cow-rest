package environment

import "os"

const (
	PROD        = "production"
	DEVELOPMENT = "dev"
	TEST        = "test"
)

func GetString(envName string, defaultValue string) string {
	if value := os.Getenv(envName); value != "" {
		return value
	}
	return defaultValue
}
