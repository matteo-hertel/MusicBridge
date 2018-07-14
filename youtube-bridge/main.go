package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := GetEnv("PORT", "3460")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}
		fmt.Fprintln(res, "Hello, World 🎉")
	})

	http.HandleFunc("/auth", redirectToAuthUrl)
	http.HandleFunc("/auth-url", authURL)
	http.HandleFunc("/auth-callback", authCallback)
	http.HandleFunc("/create-playlist", makePlaylist)
	http.HandleFunc("/search", search)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
