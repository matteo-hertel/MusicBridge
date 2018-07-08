package main

import (
	"net/http"
	"os"
)

func CheckAccessToken(token string) Adapter {
  return func(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			accessToken := req.Header.Get(token);

	if len(accessToken) == 0 {
		return 
	}
      handler.ServeHTTP(res,req)
    }
  }
}

func GetEnv(key string, fallback string) string {
	envVar := os.Getenv(key)
	if len(envVar) == 0 {
		return fallback
	}
	return envVar
}
