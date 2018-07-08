package main

import (
	"fmt"
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func main() {
	port := GetEnv("PORT", "3460")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello, World 🎉")
	})

	http.HandleFunc("/auth", redirectToAuthUrl)
	http.HandleFunc("/auth-url", authURL)
	http.HandleFunc("/auth-callback", authCallback)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		handler = adapter(handler)
	}
	return handler
}
