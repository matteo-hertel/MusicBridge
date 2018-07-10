package main

import (
	"os"
)

func GetEnv(key string, fallback string) string {
	envVar := os.Getenv(key)
	if len(envVar) == 0 {
		return fallback
	}
	return envVar
}
