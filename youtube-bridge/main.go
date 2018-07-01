package main

import (
	"fmt"
	"log"
	"net/http"

	"./youtube"

	"google.golang.org/api/youtube/v3"
)

func main() {
	port := GetEnv("PORT", "3460")
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/auth-url", authUrl)
	http.HandleFunc("/auth-callback", authCallback)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func authUrl(res http.ResponseWriter, req *http.Request) {
	config, _ := yt.GetApiConfig()
	redirectUrl := yt.GetAuthURL(config)
	http.Redirect(res, req, redirectUrl, http.StatusMovedPermanently)
}

func authCallback(res http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	config, ctx := yt.GetApiConfig()
	accessToken := yt.GetAccessToken(config, code)
	client := config.Client(ctx, accessToken)

	service, err := youtube.New(client)

	handleError(err, "Error creating YouTube client")

	data := yt.ChannelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")
	fmt.Fprintln(res, fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		data[0].Id,
		data[0].Snippet.Title,
		data[0].Statistics.ViewCount))
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
