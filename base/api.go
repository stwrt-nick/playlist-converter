package base

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
)

var (
	SpotifyClientId     = os.Getenv("SPOTIFY_CLIENT_ID")
	SpotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
)

func GetSpotifyAuthToken() (token string, err error) {

	// Load env file
	err = godotenv.Load("credentials.env")
	if err != nil {
		err = errors.New("error loading env file")
		return token, err
	}

	// Formulate Request
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token?", strings.NewReader(data.Encode()))
	SpotifyClientId = os.Getenv("SPOTIFY_CLIENT_ID")
	SpotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	authorization := SpotifyClientId + ":" + SpotifyClientSecret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authorization))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return token, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	var authResponse SpotifyOAuthResponse

	unmarshalErr := json.Unmarshal(body, &authResponse)
	if unmarshalErr != nil {
		panic(err)
	}

	return authResponse.AccessToken, err
}

func GetUsersPlaylistsSpotify(authToken string, userId string) (playlistId string, err error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/users/"+userId+"/playlists", nil)
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return playlistId, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return playlistId, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return playlistId, err
	}

	bodyString := string(body)

	stripSlash := strings.Replace(bodyString, "\\", "", -1)

	value := gjson.Get(stripSlash, "items.followers.id")
	playlistId = value.String()

	fmt.Println(playlistId)

	return playlistId, err
}
