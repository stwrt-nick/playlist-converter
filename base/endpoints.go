package base

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeConvertSpotifyToAppleEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(convertSpotifyToAppleRequest)
		v, err := s.ConvertSpotifyToApple(ctx, req.S)
		if err != nil {
			return convertSpotifyToAppleResponse{v, err.Error()}, nil
		}
		return convertSpotifyToAppleResponse{v, ""}, nil
	}
}

func makeConvertAppleToSpotifyEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(convertAppleToSpotifyRequest)
		v, err := s.ConvertAppleToSpotify(ctx, req.S)
		if err != nil {
			return convertAppleToSpotifyResponse{v, err.Error()}, nil
		}
		return convertAppleToSpotifyResponse{v, ""}, nil
	}
}
