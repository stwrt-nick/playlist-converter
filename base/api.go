package base

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

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

	var authResponse SpotifyOAuthResponse

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

	var playlistResponse SpotifyPlaylistResponse

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

	var tracksFromPlaylist GetPlaylistItems

	unmarshalErr := json.Unmarshal(body, &tracksFromPlaylist)
	if unmarshalErr != nil {
	}

	for _, items := range tracksFromPlaylist.Items {
		playlistTracks = append(playlistTracks, items.Track.ExternalIds.Isrc)
	}

	return playlistTracks, err
}

func privateKeyFromFile() (privateKey *ecdsa.PrivateKey, err error) {

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

	fmt.Println(expirationTimestamp.Unix())
	fmt.Println(now.Unix())

	token.Header["alg"] = "ES256"
	token.Header["kid"] = keyId

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func CreateApplePlaylist() (song string, err error) {
	privateKey, err := privateKeyFromFile()
	if err != nil {
		log.Fatal(err)
	}

	authToken, err := GenerateAuthToken(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.music.apple.com/v1/me/library/playlists",
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

	var applePlaylistResponse ApplePlaylistResponse

	unmarshalErr := json.Unmarshal(body, &applePlaylistResponse)
	if unmarshalErr != nil {
		panic(err)
	}

	return song, err

}
