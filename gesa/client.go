package gesa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/michimani/go-esa/internal"
)

type NewGesaClientInput struct {
	HTTPClient  *http.Client
	AccessToken string
	TeamName    string
	APIVersion  EsaAPIVersion
}

type IGesaClient interface {
	Exec(req *http.Request, r internal.IResponse) error
}

type GesaClient struct {
	client      *http.Client
	accessToken string
	teamName    string
	apiVersion  EsaAPIVersion
}

type ClientResponse struct {
	StatusCode int
	Status     string
	Response   internal.IResponse
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Duration(30) * time.Second,
}

func NewGesaClient(in *NewGesaClientInput) (*GesaClient, error) {
	if in == nil {
		return nil, fmt.Errorf("NewGesaClientInput is nil.")
	}

	if in.AccessToken == "" || in.TeamName == "" {
		return nil, fmt.Errorf("AccessToken or TeamName or both are empty.")
	}

	apiVersion := in.APIVersion
	if apiVersion.IsEmpty() {
		apiVersion = DefaultAPIVersion
	} else if !apiVersion.IsValid() {
		return nil, fmt.Errorf("Invalid esa API version.")
	}

	c := GesaClient{
		client:      defaultHTTPClient,
		accessToken: in.AccessToken,
		teamName:    in.TeamName,
		apiVersion:  apiVersion,
	}

	if in.HTTPClient != nil {
		c.client = in.HTTPClient
	}

	return &c, nil
}

func (c *GesaClient) AccessToken() string {
	if c == nil {
		return ""
	}
	return c.accessToken
}

func (c *GesaClient) TeamName() string {
	if c == nil {
		return ""
	}
	return c.teamName
}

func (c *GesaClient) CallAPI(ctx context.Context, endpoint, method string, p internal.IParameters, r internal.IResponse) error {
	req, err := c.prepare(ctx, endpoint, method, p)
	if err != nil {
		return wrapErr(err)
	}

	if n2xe, err := c.Exec(req, r); err != nil {
		return wrapErr(err)
	} else if n2xe != nil {
		return wrapWithAPIErr(n2xe)
	}

	return nil
}

var okCodes map[int]struct{} = map[int]struct{}{
	http.StatusOK:        {},
	http.StatusCreated:   {},
	http.StatusNoContent: {},
}

func (c *GesaClient) Exec(req *http.Request, r internal.IResponse) (*EsaAPIError, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if _, ok := okCodes[res.StatusCode]; !ok {
		non200err, err := resolveEsaAPIError(res)
		if err != nil {
			return nil, err
		}
		return non200err, nil
	}

	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return nil, err
	}

	r.SetRateLimitInfo(res.Header)

	return nil, nil
}

func (c *GesaClient) prepare(ctx context.Context, endpointBase, method string, p internal.IParameters) (*http.Request, error) {
	if p == nil {
		return nil, errors.New("parameter is nil")
	}

	// resolve query parameters
	endpoint, err := c.resolveEndpoint(endpointBase, p)
	if err != nil {
		return nil, err
	}

	req, err := newRequest(ctx, endpoint, method, p)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))

	return req, nil
}

func (c *GesaClient) resolveEndpoint(base string, p internal.IParameters) (string, error) {
	endpoint := p.ResolveEndpoint(base)
	return c.apiVersion.ResolveEndpoint(endpoint)
}

func newRequest(ctx context.Context, endpoint, method string, p internal.IParameters) (*http.Request, error) {
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

func resolveEsaAPIError(res *http.Response) (*EsaAPIError, error) {
	n2xe := &EsaAPIError{
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}

	cts := internal.HeaderValues("Content-Type", res.Header)
	if len(cts) == 0 {
		n2xe.Error = "Content-Type is undefined."
		return n2xe, nil
	}

	if !strings.Contains(cts[0], "application/json") {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		n2xe.Error = strings.TrimRight(string(bytes), "\n")
		return n2xe, nil
	}

	if err := json.NewDecoder(res.Body).Decode(n2xe); err != nil {
		return nil, err
	}

	// additional information for Rate Limit
	if res.StatusCode == http.StatusTooManyRequests {
		rri, err := GetRateLimitInformation(res.Header)
		if err != nil {
			return nil, err
		}

		n2xe.RateLimitInfo = rri
	}

	return n2xe, nil
}
