package base

import (
	"context"
	"errors"
	"fmt"
	"log"

	"sync"

	"github.com/stwrt-nick/playlist-converter/api"
	"github.com/stwrt-nick/playlist-converter/model"
)

type service struct{}

type Service interface {
	ConvertSpotifyToApple(ctx context.Context, req model.convertSpotifyToAppleRequest) (res model.convertSpotifyToAppleResponse, err error)
	// GetUsersPlaylistsSpotify(ctx context.Context, userID string) (res GetUsersPlaylistsSpotifyResponse, err error)
	ConvertAppleToSpotify(ctx context.Context, req model.convertAppleToSpotifyRequest) (res model.convertAppleToSpotifyResponse, err error)
	GetAppleJWTToken(ctx context.Context, req model.getAppleJWTTokenRequest) (res model.getAppleJWTTokenResponse, err error)
	// GetAppleSong(ctx context.Context, req getAppleSongRequest) (res getAppleSongResponse, err error)
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type baseService struct {
	mtx sync.RWMutex
}

func NewBaseService() Service {
	return &baseService{}
}

func (s *baseService) ConvertSpotifyToApple(ctx context.Context, req convertSpotifyToAppleRequest) (res convertSpotifyToAppleResponse, err error) {
	authToken, err := api.GetSpotifyAuthToken()
	if err != nil {
		return res, err
	}
	playlistId, err := api.GetPlaylistIdSpotify(authToken, req.Id, req.PlaylistName)
	if err != nil {
		return res, err
	}

	playlistTracksISRC, err := api.GetPlaylistTracksSpotify(authToken, playlistId)
	if err != nil {
		return res, err
	}

	status, err := api.CreateApplePlaylist(playlistTracksISRC, req.PlaylistName)
	if err != nil {
		return res, err
	}

	if status != "" {
		res = model.convertSpotifyToAppleResponse{
			Status: status,
			Err:    nil,
		}
	}

	if status == "" {
		res = convertSpotifyToAppleResponse{
			Status: "failed",
			Err:    errors.New("request failed"),
		}
	}

	return res, nil
}

func (s *baseService) ConvertAppleToSpotify(ctx context.Context, req model.convertAppleToSpotifyRequest) (res model.convertAppleToSpotifyResponse, err error) {
	return res, nil
}

func (s *baseService) GetAppleJWTToken(ctx context.Context, req model.getAppleJWTTokenRequest) (res model.getAppleJWTTokenResponse, err error) {
	fmt.Println("79")
	privateKey, err := api.privateKeyFromFile()
	if err != nil {
		log.Fatal(err)
	}
	authToken, err := api.GenerateAuthToken(privateKey)

	res.JWTToken = authToken

	return res, err
}
