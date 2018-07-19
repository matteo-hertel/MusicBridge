package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type ApiConfig struct {
	Config *oauth2.Config
}

func GetAuthURL(config *oauth2.Config) string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func GetApiConfig() (ApiConfig, error) {
	var apiConfig ApiConfig

	env := GetEnv("ENV", "development")
	b, err := ioutil.ReadFile(fmt.Sprintf("client_secret.%s.json", env))
	if err != nil {
		return apiConfig, err
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeForceSslScope)
	if err != nil {
		return apiConfig, err
	}

	apiConfig = ApiConfig{config}

	return apiConfig, nil
}

func makeService(token *oauth2.Token, ctx context.Context) (*youtube.Service, error) {
	config, err := GetApiConfig()

	if err != nil {
		return nil, err
	}
	client := config.Config.Client(ctx, token)

	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}
	return service, nil
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

func PlaylistsInsert(service *youtube.Service, playlistDefinition *BridgePlayList, part string) (*youtube.Playlist, error) {
	playlist := &youtube.Playlist{}

	properties := (map[string]string{"snippet.title": playlistDefinition.Title,
		"snippet.description":     playlistDefinition.Description,
		"snippet.tags[]":          "",
		"snippet.defaultLanguage": "",
		"status.privacyStatus":    playlistDefinition.PrivacyStatus,
	})

	resources := createResource(properties)

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
