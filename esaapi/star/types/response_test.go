package types_test

import (
	"net/http"
	"testing"

	"github.com/michimani/go-esa/esaapi/star/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_ListPostStargazersOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.ListPostStargazersOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListPostStargazersOutput{
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
			want: &types.ListPostStargazersOutput{
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
			want: &types.ListPostStargazersOutput{
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
			want: &types.ListPostStargazersOutput{
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
			want: &types.ListPostStargazersOutput{
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
			want: &types.ListPostStargazersOutput{
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
			want: &types.ListPostStargazersOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.ListPostStargazersOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_CreatePostStarOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.CreatePostStarOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.CreatePostStarOutput{
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
			want: &types.CreatePostStarOutput{
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
			want: &types.CreatePostStarOutput{
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
			want: &types.CreatePostStarOutput{
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
			want: &types.CreatePostStarOutput{
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
			want: &types.CreatePostStarOutput{
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
			want: &types.CreatePostStarOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.CreatePostStarOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_DeletePostStarOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.DeletePostStarOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.DeletePostStarOutput{
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
			want: &types.DeletePostStarOutput{
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
			want: &types.DeletePostStarOutput{
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
			want: &types.DeletePostStarOutput{
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
			want: &types.DeletePostStarOutput{
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
			want: &types.DeletePostStarOutput{
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
			want: &types.DeletePostStarOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.DeletePostStarOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_ListCommentStargazersOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.ListCommentStargazersOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListCommentStargazersOutput{
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
			want: &types.ListCommentStargazersOutput{
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
			want: &types.ListCommentStargazersOutput{
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
			want: &types.ListCommentStargazersOutput{
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
			want: &types.ListCommentStargazersOutput{
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
			want: &types.ListCommentStargazersOutput{
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
			want: &types.ListCommentStargazersOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.ListCommentStargazersOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_CreateCommentStarOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.CreateCommentStarOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.CreateCommentStarOutput{
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
			want: &types.CreateCommentStarOutput{
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
			want: &types.CreateCommentStarOutput{
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
			want: &types.CreateCommentStarOutput{
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
			want: &types.CreateCommentStarOutput{
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
			want: &types.CreateCommentStarOutput{
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
			want: &types.CreateCommentStarOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.CreateCommentStarOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_DeleteCommentStarOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.DeleteCommentStarOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.DeleteCommentStarOutput{
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
			want: &types.DeleteCommentStarOutput{
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
			want: &types.DeleteCommentStarOutput{
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
			want: &types.DeleteCommentStarOutput{
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
			want: &types.DeleteCommentStarOutput{
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
			want: &types.DeleteCommentStarOutput{
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
			want: &types.DeleteCommentStarOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.DeleteCommentStarOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}
