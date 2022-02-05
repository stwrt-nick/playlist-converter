package base

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
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

func (mw loggingMiddleware) ConvertAppleToSpotify(ctx context.Context, userID string) (res string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "convertAppleToSpotify", "id", userID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.ConvertAppleToSpotify(ctx, userID)
}

func (mw loggingMiddleware) ConvertSpotifyToApple(ctx context.Context, userID string) (res string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "convertSpotifyToApple", "id", userID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.ConvertSpotifyToApple(ctx, userID)
}

func (mw loggingMiddleware) GetSpotifyAuthToken() {
	defer func(begin time.Time) {
		mw.logger.Log("method", "convertAppleToSpotify", "took", time.Since(begin))
	}(time.Now())
}
