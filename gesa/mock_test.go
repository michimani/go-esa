package gesa_test

import (
	"io"
	"net/http"

	"github.com/michimani/go-esa/internal"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newMockClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

type mockInput struct {
	ResponseStatusCode int
	ResponseHeader     map[string][]string
	ResponseBody       io.ReadCloser
}

func newMockHTTPClient(in *mockInput) *http.Client {
	if in == nil {
		return nil
	}

	return newMockClient(func(req *http.Request) *http.Response {
		return &http.Response{
			Status:     "mock response status",
			StatusCode: in.ResponseStatusCode,
			Body:       in.ResponseBody,
			Header:     in.ResponseHeader,
		}
	})
}

type mockAPIParameter struct {
	EsaAPINil bool
}

func (mp mockAPIParameter) EsaAPIParameter() *internal.EsaAPIParameter {
	if mp.EsaAPINil {
		return nil
	}
	return &internal.EsaAPIParameter{}
}

type mockAPIResponse struct{}

func (mr mockAPIResponse) SetRateLimitInfo(h http.Header) {}
