package base

import (
	"context"
	"playlist-converter/model"
	"time"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) ConvertAppleToSpotify(ctx context.Context, req model.ConvertAppleToSpotifyRequest) (res model.ConvertAppleToSpotifyResponse, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "convertAppleToSpotify", "id", req.Id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.ConvertAppleToSpotify(ctx, req)
}

func (mw loggingMiddleware) ConvertSpotifyToApple(ctx context.Context, req model.ConvertSpotifyToAppleRequest) (res model.ConvertSpotifyToAppleResponse, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "convertSpotifyToApple", "id", req.Id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.ConvertSpotifyToApple(ctx, req)
}

func (mw loggingMiddleware) GetAppleJWTToken(ctx context.Context, req model.GetAppleJWTTokenRequest) (res model.GetAppleJWTTokenResponse, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "getAppleJWTToken", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetAppleJWTToken(ctx, req)
}

// func (mw loggingMiddleware) GetUsersPlaylistsSpotify(ctx context.Context, userID string) (res GetUsersPlaylistsSpotifyResponse, err error) {
// 	defer func(begin time.Time) {
// 		mw.logger.Log("method", "getUsersPlaylistsSpotify", "id", userID, "took", time.Since(begin), "err", err)
// 	}(time.Now())
// 	return mw.next.GetUsersPlaylistsSpotify(ctx, userID)
// }

// func (mw loggingMiddleware) GetSpotifyAuthToken() {
// 	defer func(begin time.Time) {
// 		mw.logger.Log("method", "convertAppleToSpotify", "took", time.Since(begin))
// 	}(time.Now())
// }
