package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/appengine"
)

type ApiConfig struct {
	Config *oauth2.Config
}

func GetAuthURL(config *oauth2.Config) string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func GetOauthToken(accessToken string) *oauth2.Token {
	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	return token
}

func GetApiConfig() (ApiConfig, error) {
	var apiConfig ApiConfig

	env := GetEnv("ENV", "development")
	b, err := ioutil.ReadFile(fmt.Sprintf("client_secret.%s.json", env))
	if err != nil {
		log.Println("Unable to read client secret file: %v", err)
		return apiConfig, err
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeForceSslScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: %v", err)
		return apiConfig, err
	}

	apiConfig = ApiConfig{config}

	return apiConfig, nil
}

func GetAccessToken(config *oauth2.Config, req *http.Request) (*oauth2.Token, error) {
	code := req.FormValue("code")
	ctx := appengine.NewContext(req)

	accessToken, err := config.Exchange(ctx, code)

	if err != nil {
		log.Println("Unable to retrieve token from web %v", err)
		return nil, err
	}
	return accessToken, nil
}

func makeService(res http.ResponseWriter, req *http.Request) (*youtube.Service, error) {
	var emptyService = youtube.Service{}
	accessToken, err := CheckAccessToken(req)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
		return &emptyService, err
	}
	token := GetOauthToken(accessToken)
	config, err := GetApiConfig()

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return &emptyService, err
	}
	ctx := appengine.NewContext(req)
	client := config.Config.Client(ctx, token)

	service, err := youtube.New(client)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return &emptyService, err
	}
	return service, nil
}

func CheckAccessToken(req *http.Request) (string, error) {
	accessToken := req.Header.Get("X-Youtube-Token")
	if len(accessToken) == 0 {
		err := errors.New("Missing or Invalid Token")
		return "", err
	}
	return accessToken, nil
}

func PlaylistItemInsert(service *youtube.Service, part string, resources string) (*youtube.PlaylistItem, error) {
	item := &youtube.PlaylistItem{}

	if err := json.NewDecoder(strings.NewReader(resources)).Decode(&item); err != nil {
		return item, err
	}

	call := service.PlaylistItems.Insert(part, item)
	response, err := call.Do()
	if err != nil {
		return item, err
	}

	return response, nil
}

func PlaylistsInsert(service *youtube.Service, part string, resources string) (*youtube.Playlist, error) {
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

func ChannelsListByUsername(service *youtube.Service, part string, forUsername string) ([]*youtube.Channel, error) {
	call := service.Channels.List(part)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	if err != nil {
		var yotubeChannel = youtube.Channel{}
		slice := []*youtube.Channel{&yotubeChannel}
		log.Println("Unable to get channel info", err)
		return slice, err
	}
	return response.Items, nil
}
