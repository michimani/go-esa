package gesa_test

import (
	"net/http"
	"testing"

	"github.com/michimani/go-esa/v2/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_GetRateLimitInformation(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name    string
		res     http.Header
		wantErr bool
		expect  *gesa.RateLimitInformation
	}{
		{
			name: "normal",
			res: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			wantErr: false,
			expect: &gesa.RateLimitInformation{
				Limit:     1,
				Remaining: 100,
				Reset:     &resetTimestamp,
			},
		},
		{
			name: "normal: limit is empty",
			res: http.Header{
				"X-RateLimit-Limit":     []string{},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			wantErr: false,
			expect: &gesa.RateLimitInformation{
				Limit:     0,
				Remaining: 100,
				Reset:     &resetTimestamp,
			},
		},
		{
			name: "normal: remaining is empty",
			res: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			wantErr: false,
			expect: &gesa.RateLimitInformation{
				Limit:     1,
				Remaining: 0,
				Reset:     &resetTimestamp,
			},
		},
		{
			name: "normal: reset is empty",
			res: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{},
			},
			wantErr: false,
			expect: &gesa.RateLimitInformation{
				Limit:     1,
				Remaining: 100,
				Reset:     nil,
			},
		},
		{
			name: "error: invalid rate limit limit value",
			res: http.Header{
				"X-RateLimit-Limit":     []string{"a"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "error: invalid rate limit remaining value",
			res: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"a"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "error: invalid rate limit reset value",
			res: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"a"},
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ri, err := gesa.GetRateLimitInformation(c.res)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, ri)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, ri)
		})
	}
}
