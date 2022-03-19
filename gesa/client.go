package gesa

import (
	"bytes"
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

type NewClientInput struct {
	HTTPClient  *http.Client
	AccessToken string
	APIVersion  EsaAPIVersion
	Debug       bool
}

type IClient interface {
	Exec(req *http.Request, r internal.IOutput) error
}

type Client struct {
	client      *http.Client
	accessToken string
	apiVersion  EsaAPIVersion
	debug       bool
}

type ClientResponse struct {
	StatusCode int
	Status     string
	Response   internal.IOutput
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Duration(30) * time.Second,
}

func NewClient(in *NewClientInput) (*Client, error) {
	if in == nil {
		return nil, fmt.Errorf("NewClientInput is nil.")
	}

	if in.AccessToken == "" {
		return nil, fmt.Errorf("AccessToken is empty.")
	}

	apiVersion := in.APIVersion
	if apiVersion.IsEmpty() {
		apiVersion = DefaultAPIVersion
	} else if !apiVersion.IsValid() {
		return nil, fmt.Errorf("Invalid esa API version.")
	}

	c := Client{
		client:      defaultHTTPClient,
		accessToken: in.AccessToken,
		apiVersion:  apiVersion,
	}

	if in.Debug {
		c.debug = true
	}

	if in.HTTPClient != nil {
		c.client = in.HTTPClient
	}

	return &c, nil
}

func (c *Client) AccessToken() string {
	if c == nil {
		return ""
	}
	return c.accessToken
}

func (c *Client) CallAPI(ctx context.Context, endpoint, method string, p internal.IInput, r internal.IOutput) error {
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

func (c *Client) Exec(req *http.Request, r internal.IOutput) (*EsaAPIError, error) {
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

	var tr io.Reader
	debugBuf := new(bytes.Buffer)
	if c.debug {
		tr = io.TeeReader(res.Body, debugBuf)
	} else {
		tr = res.Body
	}

	if err := json.NewDecoder(tr).Decode(r); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	if c.debug {
		fmt.Printf("------DEBUG------\nraw response\n%s\n------DEBUG END------\n", debugBuf.String())
	}

	r.SetRateLimitInfo(res.Header)

	return nil, nil
}

func (c *Client) prepare(ctx context.Context, endpointBase, method string, p internal.IInput) (*http.Request, error) {
	if p == nil {
		return nil, errors.New("parameter is nil")
	}

	eap, err := p.EsaAPIParameter()
	if err != nil {
		return nil, err
	}
	if eap == nil {
		return nil, errors.New("parameter for esa API is nil")
	}

	// resolve query parameters
	endpoint, err := c.resolveEndpoint(endpointBase, *eap)
	if err != nil {
		return nil, err
	}

	req, err := newRequest(ctx, endpoint, method, eap.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken()))

	return req, nil
}

func (c *Client) resolveEndpoint(base string, eap internal.EsaAPIParameter) (string, error) {
	endpoint := internal.ResolveEndpoint(base, eap.Path, eap.Query)
	return c.apiVersion.ResolveEndpoint(endpoint)
}

func newRequest(ctx context.Context, endpoint, method string, body io.Reader) (*http.Request, error) {
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
