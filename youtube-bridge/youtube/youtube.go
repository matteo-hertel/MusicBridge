package yt

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func GetAuthURL(config *oauth2.Config) string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func GetAccessToken(config *oauth2.Config, webToken string) *oauth2.Token {
	accessToken, err := config.Exchange(oauth2.NoContext, webToken)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return accessToken
}

func ChannelsListByUsername(service *youtube.Service, part string, forUsername string) []*youtube.Channel {
	call := service.Channels.List(part)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	return response.Items
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
func GetApiConfig() (*oauth2.Config, context.Context) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config, ctx
}
