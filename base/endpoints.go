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
	GetAppleJWTTokenEndpoint      endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		ConvertSpotifyToAppleEndpoint: makeConvertSpotifyToAppleEndpoint(s),
		ConvertAppleToSpotifyEndpoint: makeConvertAppleToSpotifyEndpoint(s),
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

func (e Endpoints) ConvertSpotifyToApple(ctx context.Context, req convertSpotifyToAppleRequest) (res convertSpotifyToAppleResponse, err error) {
	request := convertAppleToSpotifyRequest{Id: req.Id}
	response, err := e.ConvertSpotifyToAppleEndpoint(ctx, request)
	if err != nil {
		return convertSpotifyToAppleResponse{}, err
	}
	resp := response.(convertSpotifyToAppleResponse)
	return convertSpotifyToAppleResponse{
		Status: resp.Status,
		Err:    resp.Err,
	}, nil
}

func makeConvertSpotifyToAppleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(convertSpotifyToAppleRequest)
		p, e := s.ConvertSpotifyToApple(ctx, req)
		return convertSpotifyToAppleResponse{Status: p.Status, Err: e}, nil
	}
}

func makeConvertAppleToSpotifyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(convertAppleToSpotifyRequest)
		p, e := s.ConvertAppleToSpotify(ctx, req)
		return convertAppleToSpotifyResponse{Status: p.Status, Err: e}, nil
	}
}

func makeGetAppleJWTTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		p, e := s.GetAppleJWTToken(ctx)
		return AppleJWTTokenResponse{JWTToken: p.JWTToken, Err: e}, nil
	}
}
