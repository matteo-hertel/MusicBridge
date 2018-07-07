package main

import (
	"./youtube"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

func authURL(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]string)
	config, _ := yt.GetApiConfig()
	redirectUrl := yt.GetAuthURL(config)

	data["url"] = redirectUrl

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)

	if err != nil {
		handleError(err, "Error getting auth Url ")
	}

	fmt.Fprintln(res, buf.String())
}

func redirectToAuthUrl(res http.ResponseWriter, req *http.Request) {
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
