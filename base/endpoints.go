package base

import (
	"context"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Endpoints struct {
	ConvertSpotifyToAppleEndpoint endpoint.Endpoint
	ConvertAppleToSpotifyEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		ConvertSpotifyToAppleEndpoint: makeConvertSpotifyToAppleEndpoint(ctx, s),
		ConvertAppleToSpotifyEndpoint: makeConvertAppleToSpotifyEndpoint(ctx, s),
	}
}

func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	// Note that the request encoders need to modify the request URL, changing
	// the path. That's fine: we simply need to provide specific encoders for
	// each endpoint.

	return Endpoints{
		ConvertSpotifyToAppleEndpoint: httptransport.NewClient("GET", tgt, encodeConvertSpotifyToAppleRequest, decodeConvertSpotifyToAppleResponse, options...).Endpoint(),
		ConvertAppleToSpotifyEndpoint: httptransport.NewClient("GET", tgt, encodeConvertAppleToSpotifyRequest, decodeConvertAppleToSpotifyResponse, options...).Endpoint(),
	}, nil
}

func (e Endpoints) ConvertSpotifyToApple(ctx context.Context, id string) (convertSpotifyToAppleResponse, error) {
	request := convertAppleToSpotifyRequest{id: id}
	response, err := e.ConvertSpotifyToAppleEndpoint(ctx, request)
	if err != nil {
		return convertSpotifyToAppleResponse{}, err
	}
	resp := response.(convertSpotifyToAppleResponse)
	return convertSpotifyToAppleResponse{
		status: resp.status,
		err:    resp.err,
	}, nil
}

func makeConvertSpotifyToAppleEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(convertSpotifyToAppleRequest)
		v, err := s.ConvertSpotifyToApple(ctx, req.id)
		if err != nil {
			return convertSpotifyToAppleResponse{v, err.Error()}, nil
		}
		return convertSpotifyToAppleResponse{v, ""}, nil
	}
}

func makeConvertAppleToSpotifyEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(convertAppleToSpotifyRequest)
		v, err := s.ConvertAppleToSpotify(ctx, req.id)
		if err != nil {
			return convertAppleToSpotifyResponse{v, err.Error()}, nil
		}
		return convertAppleToSpotifyResponse{v, ""}, nil
	}
}
