package base

import (
	"context"
	"errors"

	"sync"
)

type service struct{}

type Service interface {
	ConvertSpotifyToApple(ctx context.Context, req convertSpotifyToAppleRequest) (res convertSpotifyToAppleResponse, err error)
	// GetUsersPlaylistsSpotify(ctx context.Context, userID string) (res GetUsersPlaylistsSpotifyResponse, err error)
	ConvertAppleToSpotify(ctx context.Context, req convertAppleToSpotifyRequest) (res convertAppleToSpotifyResponse, err error)
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

	authToken, err := GetSpotifyAuthToken()
	if err != nil {
		return res, err
	}
	playlistId, err := GetUsersPlaylistsSpotify(authToken, req.Id)
	if err != nil {
		return res, err
	}

	if req.Id != "" {
		res = convertSpotifyToAppleResponse{
			Status: playlistId,
			Err:    nil,
		}
	}

	if req.Id == "" {
		res = convertSpotifyToAppleResponse{
			Status: "failed",
			Err:    errors.New("no id found"),
		}
	}

	return res, nil
}

func (s *baseService) ConvertAppleToSpotify(ctx context.Context, req convertAppleToSpotifyRequest) (res convertAppleToSpotifyResponse, err error) {
	return res, nil
}
