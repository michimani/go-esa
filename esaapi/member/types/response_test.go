package types_test

import (
	"net/http"
	"testing"

	"github.com/michimani/go-esa/esaapi/member/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_ListMembersOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.ListMembersOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListMembersOutput{
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
			want: &types.ListMembersOutput{
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
			want: &types.ListMembersOutput{
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
			want: &types.ListMembersOutput{
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
			want: &types.ListMembersOutput{
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
			want: &types.ListMembersOutput{
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
			want: &types.ListMembersOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.ListMembersOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_DeleteMemberOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.DeleteMemberOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.DeleteMemberOutput{
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
			want: &types.DeleteMemberOutput{
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
			want: &types.DeleteMemberOutput{
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
			want: &types.DeleteMemberOutput{
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
			want: &types.DeleteMemberOutput{
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
			want: &types.DeleteMemberOutput{
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
			want: &types.DeleteMemberOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.DeleteMemberOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}
