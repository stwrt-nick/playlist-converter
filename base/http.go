package base

import (
	"context"
	"encoding/json"
	"net/http"
)

type convertSpotifyToAppleRequest struct {
	S string `json:"s"`
}

type convertSpotifyToAppleResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type convertAppleToSpotifyRequest struct {
	S string `json:"s"`
}

type convertAppleToSpotifyResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func decodeConvertSpotifyToAppleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request convertSpotifyToAppleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeConvertAppleToSpotifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request convertAppleToSpotifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
