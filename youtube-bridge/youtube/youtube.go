package yt

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type ApiConfig struct {
	Ctx    context.Context
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
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Println("Unable to read client secret file: %v", err)
		return apiConfig, err
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeForceSslScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: %v", err)
		return apiConfig, err
	}

	apiConfig = ApiConfig{ctx, config}

	return apiConfig, nil
}

func GetAccessToken(config *oauth2.Config, webToken string) (*oauth2.Token, error) {
	accessToken, err := config.Exchange(oauth2.NoContext, webToken)
	if err != nil {
		log.Println("Unable to retrieve token from web %v", err)
		return nil, err
	}
	return accessToken, nil
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
