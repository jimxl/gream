package env

import (
	"strings"
	"os"
)


var envStr = strings.ToLower(os.Getenv("GREAM_ENV"))


func Is(envName string) bool {
	return strings.TrimSpace(envName) == Env()
}

func Env() string {
	if envStr == "" {
		return "development"
	}
	return envStr
}

func IsProduction() bool {
	return Is("production")
}

func IsDevelopment() bool {
	return Is("development")
}

func IsTest() bool {
	return Is("test")
}