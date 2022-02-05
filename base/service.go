package base

import (
	"context"
)

type service struct{}

type Service interface {
	ConvertSpotifyToApple(ctx context.Context, req string) (string, error)
	ConvertAppleToSpotify(ctx context.Context, req string) (string, error)
	GetSpotifyAuthToken()
}

func ConvertSpotifyToApple(string) string {
	return "Hello"
}

func ConvertAppleToSpotify(string) string {
	return "Hello"
}
