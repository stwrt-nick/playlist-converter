package base

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"strings"
	"time"

	"github.com/joho/godotenv"
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

func GetPlaylistTracksSpotify(authToken string, userId string, playlistName string) (playlistTracks []string, err error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/users/"+userId+"/playlists", nil)
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return playlistTracks, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return playlistTracks, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return playlistTracks, err
	}

	var playlistResponse SpotifyPlaylistResponse

	unmarshalErr := json.Unmarshal(body, &playlistResponse)
	if unmarshalErr != nil {

	}

	for playlistCount := range playlistResponse.Items {
		if currentPlaylist := playlistResponse.Items[playlistCount]; currentPlaylist.Name == playlistName {
			for trackNumber, tracks := range currentPlaylist.Tracks {
				playlistTracks = append(playlistTracks, tracks.Items[trackNumber].Name)
			}
		}
	}

	return playlistTracks, err
}

func GetAppleAuthToken() {

}
