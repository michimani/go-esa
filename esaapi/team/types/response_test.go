package types_test

import (
	"net/http"
	"testing"

	"github.com/michimani/go-esa/esaapi/team/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_ListTeamsOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.ListTeamsOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 100,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "normal: limit is empty",
			h: http.Header{
				"X-RateLimit-Limit":     []string{},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     0,
					Remaining: 100,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "normal: remaining is empty",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 0,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "normal: reset is empty",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 100,
					Reset:     nil,
				},
			},
		},
		{
			name: "error: invalid rate limit limit value",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"a"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: nil,
			},
		},
		{
			name: "error: invalid rate limit remaining value",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"a"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: nil,
			},
		},
		{
			name: "error: invalid rate limit reset value",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"a"},
			},
			want: &types.ListTeamsOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.ListTeamsOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_GetTeamOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.GetTeamOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 100,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "normal: limit is empty",
			h: http.Header{
				"X-RateLimit-Limit":     []string{},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     0,
					Remaining: 100,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "normal: remaining is empty",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 0,
					Reset:     &resetTimestamp,
				},
			},
		},
		{
			name: "normal: reset is empty",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: &gesa.RateLimitInformation{
					Limit:     1,
					Remaining: 100,
					Reset:     nil,
				},
			},
		},
		{
			name: "error: invalid rate limit limit value",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"a"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: nil,
			},
		},
		{
			name: "error: invalid rate limit remaining value",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"a"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: nil,
			},
		},
		{
			name: "error: invalid rate limit reset value",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"a"},
			},
			want: &types.GetTeamOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.GetTeamOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}
