package oauth

import (
	"errors"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
)

type IOAuthProvider interface {
	OAuth(req *Request) (*Response, error)
	Init(endpoints map[string]Endpoint)
	GetEndpoint(endpoint string) (*Endpoint, error)
}

type BaseOAuth struct {
	Http      *resty.Client
	Endpoints map[string]Endpoint
}

func (s *BaseOAuth) GetEndpoint(endpoint string) (*Endpoint, error) {
	ep, exist := s.Endpoints[endpoint]
	if !exist {
		return nil, errors.New("Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *BaseOAuth) Init(endpoints map[string]Endpoint) {
	s.Endpoints = endpoints
	s.Http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *BaseOAuth) PreAuth(req *Request) (*Endpoint, error) {
	if req == nil {
		return nil, errors.New("Request Data Is Nil ")
	}

	endpoint, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return nil, err
	}

	return endpoint, nil
}