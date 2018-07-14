package main

import (
	"./youtube"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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
	http.Redirect(res, req, redirectUrl, http.StatusTemporaryRedirect)
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

func playlistsInsert(service *youtube.Service, part string, resources string) (*youtube.Playlist, error) {
	playlist := &youtube.Playlist{}

	if err := json.NewDecoder(strings.NewReader(resources)).Decode(&playlist); err != nil {
		return playlist, err
	}

	call := service.Playlists.Insert(part, playlist)
	response, err := call.Do()
	if err != nil {
		return playlist, err
	}

	return response, nil
}

func makePlaylist(res http.ResponseWriter, req *http.Request) {
	service, err := makeService(res, req)
	if err != nil {
		return
	}

	properties := (map[string]string{"snippet.title": "AutoPlayList",
		"snippet.description":     "Wad up?",
		"snippet.tags[]":          "nailedit",
		"snippet.defaultLanguage": "",
		"status.privacyStatus":    "public",
	})
	resource := createResource(properties)

	playlistsInsert(service, "snippet,status", resource)

	res.WriteHeader(http.StatusCreated)

	fmt.Fprintln(res, http.StatusText(http.StatusCreated))
}

func search(res http.ResponseWriter, req *http.Request) {
	service, err := makeService(res, req)

	if err != nil {
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

func makeService(res http.ResponseWriter, req *http.Request) (*youtube.Service, error) {
	var emptyService = youtube.Service{}
	accessToken, err := CheckAccessToken(req)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
		return &emptyService, err
	}
	token := yt.GetOauthToken(accessToken)
	config, err := yt.GetApiConfig()

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return &emptyService, err
	}
	client := config.Config.Client(config.Ctx, token)

	service, err := youtube.New(client)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return &emptyService, err
	}
	return service, nil
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
