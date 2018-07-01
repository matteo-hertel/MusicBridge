package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := GetEnv("PORT", "3460")
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/wadup", wadup)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func wadup(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Wad Up??!")
}
func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Wad Up?!")
}
