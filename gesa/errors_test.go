package gesa_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

func Test_wrapErr(t *testing.T) {
	cases := []struct {
		name    string
		err     error
		wantNil bool
	}{
		{
			name: "normal",
			err:  errors.New("error test"),
		},
		{
			name: "normal: wrapped",
			err:  gesa.ExportWrapErr(errors.New("error test")),
		},
		{
			name:    "nil",
			err:     nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)

			ge := gesa.ExportWrapErr(c.err)
			if c.wantNil {
				assert.Nil(ge)
				return
			}

			assert.NotNil(ge)
			assert.Equal("error test", ge.Error())
			assert.False(ge.OnAPI)

			un := ge.Unwrap()
			_, ok := un.(*gesa.GesaError)
			assert.False(ok)
		})
	}
}

func Test_wrapWithAPIErr(t *testing.T) {
	reset := gesa.Timestamp(1461218696)

	cases := []struct {
		name         string
		err          *gesa.EsaAPIError
		expectErrMsg string
		wantNil      bool
	}{
		{
			name:         "normal: empty",
			err:          &gesa.EsaAPIError{},
			expectErrMsg: "The esa API returned error response with a status other than 2XX series.",
		},
		{
			name: "normal: full",
			err: &gesa.EsaAPIError{
				Status:     "non 2xx error status",
				StatusCode: 500,
				Error:      "error string",
				Message:    "message string",
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     100,
					Remaining: 20,
					Reset:     &reset,
				},
			},
			expectErrMsg: strings.Join([]string{
				"The esa API returned error response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"error=\"error string\"",
				"message=\"message string\"",
				"rateLimit=100 rateLimitRemaining=20 rateLimitReset=\"1461218696\"",
			}, " "),
		},
		{
			name: "normal: partial",
			err: &gesa.EsaAPIError{
				Status:     "non 2xx error status",
				StatusCode: 500,
				Error:      "error string",
			},
			expectErrMsg: strings.Join([]string{
				"The esa API returned error response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"error=\"error string\"",
			}, " "),
		},
		{
			name:    "nil",
			err:     nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)

			ge := gesa.ExportWrapWithAPIErr(c.err)
			if c.wantNil {
				assert.Nil(ge)
				return
			}

			assert.NotNil(ge)
			assert.True(ge.OnAPI)
			assert.Equal(c.expectErrMsg, ge.Error())
		})
	}
}

func Test_EsaAPIError_Summary(t *testing.T) {
	reset := gesa.Timestamp(1461218696)

	cases := []struct {
		name         string
		err          *gesa.EsaAPIError
		expectErrMsg string
	}{
		{
			name:         "normal: empty",
			err:          &gesa.EsaAPIError{},
			expectErrMsg: "The esa API returned error response with a status other than 2XX series.",
		},
		{
			name: "normal: full",
			err: &gesa.EsaAPIError{
				Status:     "non 2xx error status",
				StatusCode: 500,
				Error:      "error string",
				Message:    "message string",
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     100,
					Remaining: 20,
					Reset:     &reset,
				},
			},
			expectErrMsg: strings.Join([]string{
				"The esa API returned error response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"error=\"error string\"",
				"message=\"message string\"",
				"rateLimit=100 rateLimitRemaining=20 rateLimitReset=\"1461218696\"",
			}, " "),
		},
		{
			name: "normal: partial",
			err: &gesa.EsaAPIError{
				Status:     "non 2xx error status",
				StatusCode: 500,
				Error:      "error string",
			},
			expectErrMsg: strings.Join([]string{
				"The esa API returned error response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"error=\"error string\"",
			}, " "),
		},
		{
			name:         "nil",
			err:          nil,
			expectErrMsg: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)

			s := c.err.Summary()
			assert.Equal(c.expectErrMsg, s)
		})
	}
}

func Test_GesaError_Error(t *testing.T) {
	cases := []struct {
		name   string
		e      *gesa.GesaError
		expect string
	}{
		{
			name:   "normal",
			e:      gesa.ExportWrapErr(errors.New("error test")),
			expect: "error test",
		},
		{
			name:   "normal: empty",
			e:      &gesa.GesaError{},
			expect: internal.ErrorUndefined,
		},
		{
			name:   "normal: nil",
			e:      nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			a := assert.New(tt)
			e := c.e.Error()
			a.Equal(c.expect, e)
		})
	}
}

func Test_Unwrap(t *testing.T) {
	cases := []struct {
		name    string
		ge      *gesa.GesaError
		wantNil bool
	}{
		{
			name: "normal",
			ge:   &gesa.GesaError{},
		},
		{
			name:    "nil",
			ge:      nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			a := assert.New(tt)
			e := c.ge.Unwrap()

			if c.wantNil {
				a.Nil(e)
				return
			}

			_, ok := e.(*gesa.GesaError)
			a.False(ok)
		})

	}

}

func Test_AccessEsaAPIErrorFields(t *testing.T) {
	reset := gesa.Timestamp(1461218696)

	n2e := &gesa.EsaAPIError{
		Status:     "non 2xx error status",
		StatusCode: 500,
		RateLimitInfo: &gesa.RateLimitInformation{
			Limit:     100,
			Remaining: 20,
			Reset:     &reset,
		},
	}

	a := assert.New(t)

	ge := gesa.ExportWrapWithAPIErr(n2e)

	a.NotNil(ge)
	a.Equal("non 2xx error status", ge.Status)
	a.Equal(500, ge.StatusCode)
	a.Equal(100, ge.RateLimitInfo.Limit)
	a.Equal(20, ge.RateLimitInfo.Remaining)
	a.Equal(reset, *ge.RateLimitInfo.Reset)
}
