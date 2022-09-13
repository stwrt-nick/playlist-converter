package base

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"playlist-converter/model"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET     /convertSpotifyToApple/             converts Spotify Playlist to Apple Playlist
	// GET     /convertAppleToSpotify/             converts Apple Playlist to Spotify Playlist

	r.Methods("GET").Path("/convertSpotifyToApple").Handler(httptransport.NewServer(
		e.ConvertSpotifyToAppleEndpoint,
		decodeConvertSpotifyToAppleRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/convertAppleToSpotify").Handler(httptransport.NewServer(
		e.ConvertAppleToSpotifyEndpoint,
		decodeConvertAppleToSpotifyRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/getAppleJWTToken").Handler(httptransport.NewServer(
		e.GetAppleJWTTokenEndpoint,
		decodeGetAppleJWTTokenRequest,
		encodeResponse,
		options...,
	))

	return r
}

func encodeConvertSpotifyToAppleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/convertSpotifyToApple")
	req.URL.Path = "/convertSpotifyToApple"
	return encodeRequest(ctx, req, request)
}

func encodeConvertAppleToSpotifyRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/convertAppleToSpotify")
	req.URL.Path = "/convertAppleToSpotify"
	return encodeRequest(ctx, req, request)
}

func encodeGetAppleJWTTokenRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/getAppleJWTToken")
	req.URL.Path = "/getAppleJWTToken"
	return encodeRequest(ctx, req, request)
}

func decodeConvertSpotifyToAppleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.ConvertSpotifyToAppleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeConvertAppleToSpotifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.ConvertAppleToSpotifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeConvertAppleToSpotifyResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response model.ConvertAppleToSpotifyResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetAppleJWTTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.GetAppleJWTTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetAppleJWTTokenResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response model.GetAppleJWTTokenResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeConvertSpotifyToAppleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response model.ConvertSpotifyToAppleResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encodeRequest likewise JSON-encodes the request to the HTTP request body.
// Don't use it directly as a transport/http.Client EncodeRequestFunc:
// profilesvc endpoints require mutating the HTTP method and request path.
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
