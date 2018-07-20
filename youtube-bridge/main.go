package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}
		fmt.Fprintln(res, fmt.Sprintf("Hello, World 🎉!\nVersion:%s", GetEnv("GAE_VERSION", "developmet")))
	})

	http.HandleFunc("/auth", redirectToAuthUrl)
	http.HandleFunc("/auth-url", authURL)
	http.HandleFunc("/auth-callback", authCallback)
	http.HandleFunc("/create-playlist", makePlaylist)
	http.HandleFunc("/add-to-playlist", addToPlaylist)
	http.HandleFunc("/search", search)
	http.HandleFunc("/bulk-search", bulkSarch)
	appengine.Main()
}
