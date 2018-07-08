package main

import (
	"./youtube"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"google.golang.org/api/youtube/v3"
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

func toJson(data map[string]string) (string, error) {

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

//func searchVideo(){
//
//	client := config.Config.Client(config.Ctx, accessToken)
//
//	service, err := youtube.New(client)
//	handleError(err, "Error creating YouTube client")
//	// Make the API call to YouTube.
//	call := service.Search.List("id,snippet").
//		Q("Amon Amarth - Raise Your Horns").
//		VideoCategoryId("10").
//		Type("video")
//
//	response, err := call.Do()
//	handleError(err, "")
//
//	// Group video, channel, and playlist results in separate lists.
//	videos := make(map[string]string)
//	channels := make(map[string]string)
//	playlists := make(map[string]string)
//	// Iterate through each item and add it to the correct list.
//	for _, item := range response.Items {
//		fmt.Println(item.Id.Kind)
//		switch item.Id.Kind {
//		case "youtube#video":
//			videos[item.Id.VideoId] = item.Snippet.Title
//		case "youtube#channel":
//			channels[item.Id.ChannelId] = item.Snippet.Title
//		case "youtube#playlist":
//			playlists[item.Id.PlaylistId] = item.Snippet.Title
//		}
//	}
//
//	printIDs(res, "Videos", videos)
//	printIDs(res, "Channels", channels)
//	printIDs(res, "Playlists", playlists)
//}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(w http.ResponseWriter, sectionName string, matches map[string]string) {
	fmt.Fprintln(w, "%v:\n", sectionName)
	for id, title := range matches {
		fmt.Fprintln(w, "[%v] %v\n", id, title)
	}
}

//	data := yt.ChannelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")
//	fmt.Fprintln(res, fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
//		"and it has %d views.",
//		data[0].Id,
//		data[0].Snippet.Title,
//		data[0].Statistics.ViewCount))

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
