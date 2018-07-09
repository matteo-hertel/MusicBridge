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
	config, err := yt.GetApiConfig()

	if err != nil {
		handleError(err, "Error getting ApiConfig")
	}

	redirectUrl := yt.GetAuthURL(config.Config)

	data["url"] = redirectUrl

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(data)

	if err != nil {
		handleError(err, "Error getting auth Url ")
	}

	fmt.Fprintln(res, buf.String())
}

func redirectToAuthUrl(res http.ResponseWriter, req *http.Request) {
	config, err := yt.GetApiConfig()
	if err != nil {
		handleError(err, "Error getting ApiConfig")
	}
	redirectUrl := yt.GetAuthURL(config.Config)
	http.Redirect(res, req, redirectUrl, http.StatusMovedPermanently)
}

func authCallback(res http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	config, err := yt.GetApiConfig()
	if err != nil {
		handleError(err, "Error getting ApiConfig")
	}

	accessToken, err := yt.GetAccessToken(config.Config, code)
	if err != nil {
		handleError(err, "Error getting accessToken")
	}
	data := map[string]string{"access_token": accessToken.AccessToken}

	buf, err := toJson(data)

	if err != nil {
		handleError(err, "Error compressing data")
	}

	fmt.Fprintln(res, buf)
}

func search(res http.ResponseWriter, req *http.Request) {

	accessToken := req.Header.Get("X-Yutube-Token")
	token := yt.GetOauthToken(accessToken)
	config, err := yt.GetApiConfig()

	if err != nil {
		handleError(err, "Error getting ApiConfig")
	}
	client := config.Config.Client(config.Ctx, token)

	service, err := youtube.New(client)
	handleError(err, "Error creating YouTube client")
	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q("Amon Amarth - Raise Your Horns").
		VideoCategoryId("10").
		Type("video")

	response, err := call.Do()
	handleError(err, "")

	videos := make(map[string]string)
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	buf, err := toJson(videos)

	if err != nil {
		handleError(err, "Error compressing data")
	}

	fmt.Fprintln(res, buf)
}

func toJson(data map[string]string) (string, error) {

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
