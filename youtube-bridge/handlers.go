package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BridgePlayList struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	PrivacyStatus string `json:"privacyStatus"`
}

type BridgePlayListItem struct {
	PlaylistId string `json:"playlistId"`
	VideoId    string `json:"videoId"`
	Position   string `json:"posittion"`
}

func authURL(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]string)
	config, err := GetApiConfig()

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	redirectUrl := GetAuthURL(config.Config)

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
	config, err := GetApiConfig()
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	redirectUrl := GetAuthURL(config.Config)
	http.Redirect(res, req, redirectUrl, http.StatusTemporaryRedirect)
}

func authCallback(res http.ResponseWriter, req *http.Request) {
	config, err := GetApiConfig()
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	accessToken, err := GetAccessToken(config.Config, req)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	buf, err := toJson(accessToken)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	fmt.Fprintln(res, buf)
}

func makePlaylist(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
	}

	var data BridgePlayList

	err = json.Unmarshal(body, &data)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	}

	service, err := makeService(res, req)
	if err != nil {
		return
	}

	properties := (map[string]string{"snippet.title": data.Title,
		"snippet.description":     data.Description,
		"snippet.tags[]":          "",
		"snippet.defaultLanguage": "",
		"status.privacyStatus":    data.PrivacyStatus,
	})

	resource := createResource(properties)

	playlist, err := PlaylistsInsert(service, "snippet,status", resource)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	buf, err := toJson(playlist)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	res.WriteHeader(http.StatusCreated)
	fmt.Fprintln(res, buf)
}

func addToPlaylist(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
	}

	var data BridgePlayListItem

	err = json.Unmarshal(body, &data)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	}

	service, err := makeService(res, req)
	if err != nil {
		return
	}

	properties := (map[string]string{"snippet.playlistId": data.PlaylistId,
		"snippet.resourceId.kind":    "youtube#video",
		"snippet.resourceId.videoId": data.VideoId,
		"snippet.position":           data.Position,
	})
	resource := createResource(properties)

	item, err := PlaylistItemInsert(service, "snippet", resource)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	buf, err := toJson(item)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	res.WriteHeader(http.StatusCreated)
	fmt.Fprintln(res, buf)
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

func handleHttpError(res http.ResponseWriter, e StatusError) {
	fmt.Println(e.Error())
	http.Error(res, e.Error(), e.Status())
}
