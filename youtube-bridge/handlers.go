package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/appengine"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
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

type BridgeSong struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func authURL(res http.ResponseWriter, req *http.Request) {
	config, err := GetApiConfig()
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	redirectUrl := GetAuthURL(config.Config)

	data := map[string]string{
		"url": redirectUrl,
	}

	buf, err := toJson(data)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	fmt.Fprintln(res, buf)
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
		return
	}

	accessToken, err := CheckAccessToken(req)
	token := GetOauthToken(accessToken)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusUnauthorized, err})
		return
	}
	ctx := appengine.NewContext(req)
	service, err := makeService(token, ctx)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	playlist, err := PlaylistsInsert(service, &data, "snippet,status")
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
	accessToken, err := CheckAccessToken(req)
	token := GetOauthToken(accessToken)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusUnauthorized, err})
		return
	}
	ctx := appengine.NewContext(req)
	service, err := makeService(token, ctx)

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
func tmp(data BridgeSong, ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- fmt.Sprintf("%s: %s", data.Artist, data.Title)
}

func bulkSarch(res http.ResponseWriter, req *http.Request) {
	accessToken, err := CheckAccessToken(req)
	token := GetOauthToken(accessToken)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusUnauthorized, err})
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
	}

	var data []BridgeSong

	err = json.Unmarshal(body, &data)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	}

	ctx := appengine.NewContext(req)
	service, err := makeService(token, ctx)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	ch := make(chan map[string]string)
	var wg sync.WaitGroup
	wg.Add(len(data))

	for _, song := range data {
		go func(song BridgeSong) {
			defer wg.Done()
			items, err := Search(service, &song)

			if err != nil {
				return
			}
			ch <- items
		}(song)
	}

	response := make([]map[string]string, len(data))
	go func() {
		for song := range ch {
			fmt.Println(song)
			response = append(response, song)
		}
	}()

	wg.Wait()
	buf, err := toJson(response)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	fmt.Fprintln(res, buf)
	//	accessToken, err := CheckAccessToken(req)
	//	token := GetOauthToken(accessToken)
	//	if err != nil {
	//		handleHttpError(res, StatusError{http.StatusUnauthorized, err})
	//		return
	//	}
	//	ctx := appengine.NewContext(req)
	//	service, err := makeService(token, ctx)
	//
	//	if err != nil {
	//		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	//		return
	//	}
	//	items, err := Search(service, &data)
	//	if err != nil {
	//		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	//		return
	//	}
	//
	//	buf, err := toJson(items)
	//
	//	if err != nil {
	//		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	//		return
	//	}
	//
	//	fmt.Fprintln(res, buf)
}

func search(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusBadRequest, err})
	}

	var data BridgeSong

	err = json.Unmarshal(body, &data)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
	}

	fmt.Println(data)
	return
	accessToken, err := CheckAccessToken(req)
	token := GetOauthToken(accessToken)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusUnauthorized, err})
		return
	}
	ctx := appengine.NewContext(req)
	service, err := makeService(token, ctx)

	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}
	items, err := Search(service, &data)
	if err != nil {
		handleHttpError(res, StatusError{http.StatusInternalServerError, err})
		return
	}

	buf, err := toJson(items)

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

func GetAccessToken(config *oauth2.Config, req *http.Request) (*oauth2.Token, error) {
	code := req.FormValue("code")
	ctx := appengine.NewContext(req)

	accessToken, err := config.Exchange(ctx, code)

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func CheckAccessToken(req *http.Request) (string, error) {
	accessToken := req.Header.Get("X-Youtube-Token")
	if len(accessToken) == 0 {
		err := errors.New("Missing or Invalid Token")
		return "", err
	}
	return accessToken, nil
}
func GetOauthToken(accessToken string) *oauth2.Token {
	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	return token
}
