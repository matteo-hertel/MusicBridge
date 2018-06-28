package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := GetEnv("PORT", "3460")
	fmt.Println(port)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("wad up")
}
