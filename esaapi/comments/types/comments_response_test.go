package types_test

import (
	"net/http"
	"testing"

	"github.com/michimani/go-esa/esaapi/comments/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_ListPostCommentsOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.ListPostCommentsOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListPostCommentsOutput{
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
			want: &types.ListPostCommentsOutput{
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
			want: &types.ListPostCommentsOutput{
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
			want: &types.ListPostCommentsOutput{
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
			want: &types.ListPostCommentsOutput{
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
			want: &types.ListPostCommentsOutput{
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
			want: &types.ListPostCommentsOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.ListPostCommentsOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_GetCommentOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.GetCommentOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.GetCommentOutput{
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
			want: &types.GetCommentOutput{
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
			want: &types.GetCommentOutput{
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
			want: &types.GetCommentOutput{
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
			want: &types.GetCommentOutput{
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
			want: &types.GetCommentOutput{
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
			want: &types.GetCommentOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.GetCommentOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_CreateCommentOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.CreateCommentOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.CreateCommentOutput{
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
			want: &types.CreateCommentOutput{
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
			want: &types.CreateCommentOutput{
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
			want: &types.CreateCommentOutput{
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
			want: &types.CreateCommentOutput{
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
			want: &types.CreateCommentOutput{
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
			want: &types.CreateCommentOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.CreateCommentOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_UpdateCommentOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.UpdateCommentOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.UpdateCommentOutput{
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
			want: &types.UpdateCommentOutput{
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
			want: &types.UpdateCommentOutput{
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
			want: &types.UpdateCommentOutput{
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
			want: &types.UpdateCommentOutput{
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
			want: &types.UpdateCommentOutput{
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
			want: &types.UpdateCommentOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.UpdateCommentOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_DeleteCommentOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.DeleteCommentOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.DeleteCommentOutput{
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
			want: &types.DeleteCommentOutput{
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
			want: &types.DeleteCommentOutput{
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
			want: &types.DeleteCommentOutput{
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
			want: &types.DeleteCommentOutput{
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
			want: &types.DeleteCommentOutput{
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
			want: &types.DeleteCommentOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.DeleteCommentOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}

func Test_ListTeamCommentsOutput_SetRateLimitInfo(t *testing.T) {
	resetTimestamp := gesa.Timestamp(100000000)

	cases := []struct {
		name string
		h    http.Header
		want *types.ListTeamCommentsOutput
	}{
		{
			name: "normal",
			h: http.Header{
				"X-RateLimit-Limit":     []string{"1"},
				"X-RateLimit-Remaining": []string{"100"},
				"X-RateLimit-Reset":     []string{"100000000"},
			},
			want: &types.ListTeamCommentsOutput{
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
			want: &types.ListTeamCommentsOutput{
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
			want: &types.ListTeamCommentsOutput{
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
			want: &types.ListTeamCommentsOutput{
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
			want: &types.ListTeamCommentsOutput{
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
			want: &types.ListTeamCommentsOutput{
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
			want: &types.ListTeamCommentsOutput{
				RateLimitInfo: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			res := &types.ListTeamCommentsOutput{}
			res.SetRateLimitInfo(c.h)

			asst.Equal(c.want.RateLimitInfo, res.RateLimitInfo)
		})
	}
}
