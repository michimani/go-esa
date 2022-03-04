package gesa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type NewGesaClientInput struct {
	HTTPClient  *http.Client
	AccessToken string
	TeamName    string
}

type IGesaClient interface {
	Exec(req *http.Request, r IResponse) error
}

type GesaClient struct {
	client      *http.Client
	accessToken string
	teamName    string
}

type ClientResponse struct {
	StatusCode int
	Status     string
	Response   IResponse
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Duration(30) * time.Second,
}

func NewGesaClient(in *NewGesaClientInput) (*GesaClient, error) {
	if in == nil {
		return nil, fmt.Errorf("NewGesaClientInput is nil.")
	}

	c := GesaClient{
		client:      defaultHTTPClient,
		accessToken: in.AccessToken,
		teamName:    in.TeamName,
	}

	if in.HTTPClient != nil {
		c.client = in.HTTPClient
	}

	return &c, nil
}

func (c *GesaClient) CallAPI(ctx context.Context, endpoint, method string, p IParameters, r IResponse) error {
	req, err := c.prepare(ctx, endpoint, method, p)
	if err != nil {
		return err
	}

	if err := c.Exec(req, r); err != nil {
		return err
	}

	return nil
}

func (c *GesaClient) Exec(req *http.Request, r IResponse) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return err
	}

	return nil
}

func (c *GesaClient) prepare(ctx context.Context, endpointBase, method string, p IParameters) (*http.Request, error) {
	if p == nil {
		return nil, errors.New("parameter is nil")
	}

	endpoint := c.resolveEndpoint(endpointBase, p)
	req, err := newRequest(ctx, endpoint, method, p)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	return req, nil
}

func (c *GesaClient) resolveEndpoint(base string, p IParameters) string {
	return base
}

func newRequest(ctx context.Context, endpoint, method string, p IParameters) (*http.Request, error) {
	body, err := p.Body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	return req, nil
}
