package main

import (
	"./youtube"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

func authURL(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]string)
	config, err := yt.GetApiConfig()

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	redirectUrl := yt.GetAuthURL(config.Config)

	data["url"] = redirectUrl

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(data)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	fmt.Fprintln(res, buf.String())
}

func redirectToAuthUrl(res http.ResponseWriter, req *http.Request) {
	config, err := yt.GetApiConfig()
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	redirectUrl := yt.GetAuthURL(config.Config)
	http.Redirect(res, req, redirectUrl, http.StatusMovedPermanently)
}

func authCallback(res http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	config, err := yt.GetApiConfig()
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	accessToken, err := yt.GetAccessToken(config.Config, code)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	data := map[string]string{"access_token": accessToken.AccessToken}

	buf, err := toJson(data)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	fmt.Fprintln(res, buf)
}

func search(res http.ResponseWriter, req *http.Request) {

	accessToken, err := CheckAccessToken(req)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
		return
	}
	token := yt.GetOauthToken(accessToken)
	config, err := yt.GetApiConfig()

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	client := config.Config.Client(config.Ctx, token)

	service, err := youtube.New(client)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q("Amon Amarth - Raise Your Horns").
		VideoCategoryId("10").
		Type("video")

	response, err := call.Do()
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	videos := make(map[string]string)
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	buf, err := toJson(videos)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
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

func handleHttpError(res http.ResponseWriter, e StatusError) {
	fmt.Println(e.Error())
	http.Error(res, e.Error(), e.Status())
}

func CheckAccessToken(req *http.Request) (string, error) {
	accessToken := req.Header.Get("X-Youtube-Token")
	if len(accessToken) == 0 {
		err := errors.New("Missing or Invalid Token")
		return "", err
	}
	return accessToken, nil
}
