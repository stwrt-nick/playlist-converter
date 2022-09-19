package base

import (
	"context"
	"errors"
	"log"

	"sync"

	"playlist-converter/api"
	"playlist-converter/model"
)

type service struct{}

type Service interface {
	ConvertSpotifyToApple(ctx context.Context, req model.ConvertSpotifyToAppleRequest) (res model.ConvertSpotifyToAppleResponse, err error)
	// GetUsersPlaylistsSpotify(ctx context.Context, userID string) (res GetUsersPlaylistsSpotifyResponse, err error)
	ConvertAppleToSpotify(ctx context.Context, req model.ConvertAppleToSpotifyRequest) (res model.ConvertAppleToSpotifyResponse, err error)
	GetAppleJWTToken(ctx context.Context, req model.GetAppleJWTTokenRequest) (res model.GetAppleJWTTokenResponse, err error)
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

func (s *baseService) ConvertSpotifyToApple(ctx context.Context, req model.ConvertSpotifyToAppleRequest) (res model.ConvertSpotifyToAppleResponse, err error) {
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

	if status == "" {
		res = model.ConvertSpotifyToAppleResponse{
			Status: "failed",
			Err:    errors.New("failed to create apple playlist"),
		}
	}

	return res, nil
}

func (s *baseService) ConvertAppleToSpotify(ctx context.Context, req model.ConvertAppleToSpotifyRequest) (res model.ConvertAppleToSpotifyResponse, err error) {
	return res, nil
}

func (s *baseService) GetAppleJWTToken(ctx context.Context, req model.GetAppleJWTTokenRequest) (res model.GetAppleJWTTokenResponse, err error) {
	privateKey, err := api.PrivateKeyFromFile()
	if err != nil {
		log.Fatal(err)
	}
	authToken, err := api.GenerateAuthToken(privateKey)

	res.JWTToken = authToken

	return res, err
}
