package api

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"playlist-converter/model"

	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var (
	SpotifyClientId     = os.Getenv("SPOTIFY_CLIENT_ID")
	SpotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	// keyFile             = os.Getenv("KEY_FILE")
	// issuerID            = os.Getenv("TEAM_ID")
	// keyId               = os.Getenv("KEY_ID")
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

	var authResponse model.SpotifyOAuthResponse

	unmarshalErr := json.Unmarshal(body, &authResponse)
	if unmarshalErr != nil {
		panic(err)
	}

	return authResponse.AccessToken, err
}

func GetPlaylistIdSpotify(authToken string, userId string, playlistName string) (playlistId string, err error) {

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

	var playlistResponse model.SpotifyPlaylistResponse

	unmarshalErr := json.Unmarshal(body, &playlistResponse)
	if unmarshalErr != nil {
		return playlistId, err
	}

	for playlistCount := range playlistResponse.Items {
		if currentPlaylist := playlistResponse.Items[playlistCount]; currentPlaylist.Name == playlistName {
			playlistId = currentPlaylist.Id
		}
	}

	return playlistId, err
}

func GetPlaylistTracksSpotify(authToken string, playlistId string) (playlistTracks []string, err error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+playlistId+"/tracks", nil)
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

	var tracksFromPlaylist model.GetPlaylistItems

	unmarshalErr := json.Unmarshal(body, &tracksFromPlaylist)
	if unmarshalErr != nil {
		return playlistTracks, err
	}

	for _, items := range tracksFromPlaylist.Items {
		playlistTracks = append(playlistTracks, items.Track.ExternalIds.Isrc)
	}

	return playlistTracks, err
}

func PrivateKeyFromFile() (privateKey *ecdsa.PrivateKey, err error) {

	err = godotenv.Load("credentials.env")
	if err != nil {
		err = errors.New("error loading env file")
		return privateKey, err
	}
	keyFile := os.Getenv("KEY_FILE")

	bytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("AuthKey must be a valid .p8 PEM file")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, errors.New("AuthKey must be of type ecdsa.PrivateKey")
	}

}

func GenerateAuthToken(privateKey *ecdsa.PrivateKey) (JWTToken string, err error) {
	err = godotenv.Load("credentials.env")
	if err != nil {
		err = errors.New("error loading env file")
		return JWTToken, err
	}
	issuerID := os.Getenv("TEAM_ID")
	keyId := os.Getenv("KEY_ID")

	expirationTimestamp := time.Now().Add(15 * time.Hour)
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": issuerID,
		"iat": now.Unix(),
		"exp": expirationTimestamp.Unix(),
	})

	token.Header["alg"] = "ES256"
	token.Header["kid"] = keyId

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func CreateApplePlaylist(playlistTracksISRC []string, playlistName string) (status string, err error) {
	var playlistSongIDs []string

	i := 0
	for i < 3 {
		songId, err := GetAppleSongIDByISRC(playlistTracksISRC[i])
		if err != nil {
			log.Fatal(err)
		}
		playlistSongIDs = append(playlistSongIDs, songId)
		i++
	}

	privateKey, err := PrivateKeyFromFile()
	if err != nil {
		log.Fatal(err)
	}

	authToken, err := GenerateAuthToken(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	playlistAttributes := model.PlaylistAttributes{
		Description: "",
		Name:        playlistName,
	}

	playlistRelationships := []model.PlaylistData{
		model.PlaylistData{
			Id:   playlistSongIDs[0],
			Type: "song",
		},
	}

	jsonBody := model.ApplePlaylistRequest{
		Attributes: playlistAttributes,
		Data:       playlistRelationships,
	}

	body, _ := json.Marshal(jsonBody)

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.music.apple.com/v1/me/library/playlists",
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	// req.Header.Set("User-Agent", "App Store Connect Client")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var applePlaylistResponse model.ApplePlaylistResponse

	unmarshalErr := json.Unmarshal(body, &applePlaylistResponse)
	if unmarshalErr != nil {
		log.Fatal(err)
	}

	return status, err

}

func GetAppleSongIDByISRC(isrc string) (songId string, err error) {
	privateKey, err := PrivateKeyFromFile()
	if err != nil {
		log.Fatal(err)
	}

	authToken, err := GenerateAuthToken(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.music.apple.com/v1/catalog/us/songs?filter[isrc]="+isrc,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	// req.Header.Set("User-Agent", "App Store Connect Client")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var appleSongIdResponse model.GetAppleSongIDByISRCResponse

	unmarshalErr := json.Unmarshal(body, &appleSongIdResponse)
	if unmarshalErr != nil {
		return songId, err
	}

	songId = appleSongIdResponse.Data[0].Id

	return songId, err
}
