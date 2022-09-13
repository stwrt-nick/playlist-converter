package base

import (
	"context"
	"net/url"
	"playlist-converter/model"
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
		GetAppleJWTTokenEndpoint:      makeGetAppleJWTTokenEndpoint(s),
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
		GetAppleJWTTokenEndpoint:      httptransport.NewClient("GET", tgt, encodeGetAppleJWTTokenRequest, decodeGetAppleJWTTokenResponse, options...).Endpoint(),
	}, nil
}

func (e Endpoints) ConvertSpotifyToApple(ctx context.Context, req model.ConvertSpotifyToAppleRequest) (res model.ConvertSpotifyToAppleResponse, err error) {
	request := model.ConvertAppleToSpotifyRequest{Id: req.Id}
	response, err := e.ConvertSpotifyToAppleEndpoint(ctx, request)
	if err != nil {
		return model.ConvertSpotifyToAppleResponse{}, err
	}
	resp := response.(model.ConvertSpotifyToAppleResponse)
	return model.ConvertSpotifyToAppleResponse{
		Status: resp.Status,
		Err:    resp.Err,
	}, nil
}

func makeConvertSpotifyToAppleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.ConvertSpotifyToAppleRequest)
		p, e := s.ConvertSpotifyToApple(ctx, req)
		return model.ConvertSpotifyToAppleResponse{Status: p.Status, Err: e}, nil
	}
}

func makeConvertAppleToSpotifyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.ConvertAppleToSpotifyRequest)
		p, e := s.ConvertAppleToSpotify(ctx, req)
		return model.ConvertAppleToSpotifyResponse{Status: p.Status, Err: e}, nil
	}
}

func makeGetAppleJWTTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.GetAppleJWTTokenRequest)
		p, e := s.GetAppleJWTToken(ctx, req)
		return model.GetAppleJWTTokenResponse{JWTToken: p.JWTToken, Err: e}, nil
	}
}
