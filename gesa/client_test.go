package gesa_test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

type testParameter struct {
	BodyResErr bool
}

func (mp testParameter) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	return &internal.EsaAPIParameter{}, nil
}

func Test_NewGesaClient(t *testing.T) {
	cases := []struct {
		name    string
		in      *gesa.NewGesaClientInput
		wantErr bool
	}{
		{
			name: "ok",
			in: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
		},
		{
			name: "ok: with http client",
			in: &gesa.NewGesaClientInput{
				HTTPClient: &http.Client{
					Timeout: time.Duration(300) * time.Second,
				},
				AccessToken: "test-token",
			},
		},
		{
			name: "ok: api version",
			in: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
				APIVersion:  gesa.DefaultAPIVersion,
			},
		},
		{
			name: "ok: debug",
			in: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
				Debug:       true,
			},
		},
		{
			name:    "ng: empty parameters",
			in:      &gesa.NewGesaClientInput{},
			wantErr: true,
		},
		{
			name:    "ng: empty parameter",
			in:      &gesa.NewGesaClientInput{},
			wantErr: true,
		},
		{
			name: "ng: invalid api version",
			in: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
				APIVersion:  gesa.EsaAPIVersion("invalid version"),
			},
			wantErr: true,
		},
		{
			name:    "ng: nil",
			in:      nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			client, err := gesa.NewGesaClient(c.in)
			if c.wantErr {
				asst.NotNil(err)
				asst.Nil(client)
				return
			}

			asst.NotNil(client)
			asst.Equal(c.in.AccessToken, client.AccessToken())
		})
	}
}

func Test_GesaClient_AccessToken(t *testing.T) {
	okClient, _ := gesa.NewGesaClient(&gesa.NewGesaClientInput{AccessToken: "test-token"})
	cases := []struct {
		name   string
		client *gesa.GesaClient
		expect string
	}{
		{"not nil", okClient, "test-token"},
		{"nil", nil, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			a := c.client.AccessToken()
			asst.Equal(c.expect, a)
		})
	}
}

func Test_CallAPI(t *testing.T) {
	cases := []struct {
		name        string
		mockInput   *mockInput
		clientInput *gesa.NewGesaClientInput
		endpoint    string
		method      string
		params      internal.IParameters
		response    internal.IResponse
		wantErr     bool
	}{
		{
			name: "ok",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
			endpoint: "test-endpoint",
			method:   http.MethodGet,
			params:   &mockAPIParameter{},
			response: &mockAPIResponse{},
			wantErr:  false,
		},
		{
			name: "ok: debug",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
				Debug:       true,
			},
			endpoint: "test-endpoint",
			method:   http.MethodGet,
			params:   &mockAPIParameter{},
			response: &mockAPIResponse{},
			wantErr:  false,
		},
		{
			name: "error: parameter is nil",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
			endpoint: "test-endpoint",
			method:   http.MethodGet,
			params:   nil,
			response: &mockAPIResponse{},
			wantErr:  true,
		},
		{
			name: "error: required parameters are empty",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
			endpoint: "test-endpoint",
			method:   http.MethodGet,
			params:   &mockAPIParameter{EsaAPINil: true},
			response: &mockAPIResponse{},
			wantErr:  true,
		},
		{
			name: "error: not 200 response",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
			endpoint: "test-endpoint",
			method:   http.MethodGet,
			params:   &mockAPIParameter{},
			response: &mockAPIResponse{},
			wantErr:  true,
		},
		{
			name: "error: failed to decode json",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`///`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
			endpoint: "test-endpoint",
			method:   http.MethodGet,
			params:   &mockAPIParameter{},
			response: &mockAPIResponse{},
			wantErr:  true,
		},
		{
			name: "error: invalid method",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			clientInput: &gesa.NewGesaClientInput{
				AccessToken: "test-token",
			},
			endpoint: "test-endpoint",
			method:   "invalid method",
			params:   &mockAPIParameter{},
			response: &mockAPIResponse{},
			wantErr:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			mockClient := newMockHTTPClient(c.mockInput)
			in := c.clientInput
			in.HTTPClient = mockClient
			client, _ := gesa.NewGesaClient(in)

			err := client.CallAPI(context.Background(), c.endpoint, c.method, c.params, c.response)
			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.Nil(tt, err)
		})
	}
}

func Test_Exec(t *testing.T) {
	nonErrReq, _ := http.NewRequestWithContext(context.TODO(), "GET", "https://example.com", nil)
	errReq := &http.Request{Method: "invalid method"}

	cases := []struct {
		name            string
		mockInput       *mockInput
		req             *http.Request
		wantErr         bool
		wantEsaAPIError bool
	}{
		{
			name: "ok",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			req:             nonErrReq,
			wantErr:         false,
			wantEsaAPIError: false,
		},
		{
			name: "error: not 200 error",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			req:             nonErrReq,
			wantErr:         false,
			wantEsaAPIError: true,
		},
		{
			name: "error: cannot resolve 200 error",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseHeader: map[string][]string{
					"Content-Type": {"application/json;charset=UTF-8"},
				},
				ResponseBody: io.NopCloser(strings.NewReader(`///`)),
			},
			req:             nonErrReq,
			wantErr:         true,
			wantEsaAPIError: false,
		},
		{
			name: "error: http.Client.Do error",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			req:             errReq,
			wantErr:         true,
			wantEsaAPIError: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			mockClient := newMockHTTPClient(c.mockInput)
			client, _ := gesa.NewGesaClient(&gesa.NewGesaClientInput{
				HTTPClient:  mockClient,
				AccessToken: "test-token",
			})

			esaAPIError, err := client.Exec(c.req, &mockAPIResponse{})

			if c.wantErr {
				assert.Nil(tt, esaAPIError)
				assert.Error(tt, err)
				return
			}

			if c.wantEsaAPIError {
				assert.Nil(tt, err)
				assert.NotNil(tt, esaAPIError)
				return
			}

			assert.Nil(tt, err)
			assert.Nil(tt, esaAPIError)
		})
	}
}

func Test_newRequest(t *testing.T) {
	cases := []struct {
		name     string
		method   string
		endpoint string
		p        testParameter
		wantErr  bool
		expect   *http.Request
	}{
		{
			name:     "normal: GET",
			method:   "GET",
			endpoint: "endpoint",
			p:        testParameter{},
			wantErr:  false,
			expect: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "endpoint"},
				Header: http.Header{
					"Content-Type": []string{"application/json;charset=UTF-8"},
				},
			},
		},
		{
			name:     "normal: POST",
			method:   "POST",
			endpoint: "endpoint",
			p:        testParameter{},
			wantErr:  false,
			expect: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "endpoint"},
				Header: http.Header{
					"Content-Type": []string{"application/json;charset=UTF-8"},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			eap, err := c.p.EsaAPIParameter()
			if c.wantErr {
				asst.Error(err)
				asst.Nil(eap)
				return
			}

			r, err := gesa.ExportNewRequest(context.Background(), c.endpoint, c.method, eap.Body)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(r)
				return
			}

			asst.Equal(c.expect.Method, r.Method)
			asst.Equal(c.expect.URL, r.URL)
			asst.Equal(c.expect.Header, r.Header)
		})
	}
}

func Test_resolveEsaAPIError(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name             string
		res              *http.Response
		hasRateLimitInfo bool
		wantErr          bool
		expect           gesa.EsaAPIError
	}{
		{
			name: "normal: no rate limit error",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header: map[string][]string{
					"Content-Type": {"application/json;charset=UTF-8"},
				},
				Body: io.NopCloser(strings.NewReader(`{"message": "error"}`)),
			},
			wantErr: false,
			expect: gesa.EsaAPIError{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "normal: content-type is text/plain",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header: map[string][]string{
					"Content-Type": {"text/plain"},
				},
				Body: io.NopCloser(strings.NewReader("text error message")),
			},
			wantErr: false,
			expect: gesa.EsaAPIError{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "normal: content-type is empty",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header:     map[string][]string{},
			},
			wantErr: false,
			expect: gesa.EsaAPIError{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "normal: rate limit error",
			res: &http.Response{
				Status:     "Too Many Requests",
				StatusCode: http.StatusTooManyRequests,
				Header: map[string][]string{
					"Content-Type":          {"application/json;charset=UTF-8"},
					"X-RateLimit-Limit":     {"1"},
					"X-RateLimit-Remaining": {"2"},
					"X-RateLimit-Reset":     {"100000000"},
				},
				Body: io.NopCloser(strings.NewReader(`{"message": "error"}`)),
			},
			hasRateLimitInfo: true,
			wantErr:          false,
			expect: gesa.EsaAPIError{
				Status:     "Too Many Requests",
				StatusCode: http.StatusTooManyRequests,
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 2,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "error: failed to decode json",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header: map[string][]string{
					"Content-Type": {"application/json;charset=UTF-8"},
				},
				Body: io.NopCloser(strings.NewReader(`///`)),
			},
			wantErr: true,
		},
		{
			name: "error: on getting rate limit information",
			res: &http.Response{
				Status:     "Too Many Requests",
				StatusCode: http.StatusTooManyRequests,
				Header: map[string][]string{
					"Content-Type":          {"application/json;charset=UTF-8"},
					"X-RateLimit-Limit":     {"1"},
					"X-RateLimit-Remaining": {"xxxx"},
					"X-RateLimit-Reset":     {"100000000"},
				},
				Body: io.NopCloser(strings.NewReader(`{"message": "error"}`)),
			},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			e, err := gesa.ExportResolveEsaAPIError(c.res)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, e)
				return
			}

			assert.NoError(tt, err)

			assert.Equal(tt, c.expect.Status, e.Status)
			assert.Equal(tt, c.expect.StatusCode, e.StatusCode)
			if c.hasRateLimitInfo {
				assert.NotNil(tt, e.RateLimitInfo)
			}
		})
	}
}
