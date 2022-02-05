package base

import (
	"context"
	"errors"
)

type service struct{}

type Service interface {
	ConvertSpotifyToApple(ctx context.Context, req string) (string, error)
	GetUsersPlaylistsSpotify(ctx context.Context, userID string) (res GetUsersPlaylistsSpotifyResponse, err error)
	ConvertAppleToSpotify(ctx context.Context, req string) (string, error)
	GetSpotifyAuthToken()
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

func ConvertSpotifyToApple(string) string {
	return "Hello"
}

func ConvertAppleToSpotify(string) string {
	return "Hello"
}
