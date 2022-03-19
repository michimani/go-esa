package types_test

import (
	"strings"
	"testing"

	"github.com/michimani/go-esa/esaapi/post/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ListPostsSort_IsValid(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect bool
	}{
		{"ok", "updated", true},
		{"ok", "created", true},
		{"ok", "number", true},
		{"ok", "stars", true},
		{"ok", "watches", true},
		{"ok", "comments", true},
		{"ok", "best_match", true},
		{"ng", "unknown", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			asst.Equal(c.expect, types.ListPostsSort(c.s).IsValid())
		})
	}
}

func Test_ListPostsOrder_IsValid(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect bool
	}{
		{"ok", "asc", true},
		{"ok", "desc", true},
		{"ng", "unknown", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			asst.Equal(c.expect, types.ListPostsOrder(c.s).IsValid())
		})
	}
}

func Test_ListPostsInput_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListPostsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListPostsInput{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListPostsInput{}, 0, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			i, b := c.p.PageValue()
			asst.Equal(c.expectInt, i)
			asst.Equal(c.expectBool, b)
		})
	}
}

func Test_ListPostsInput_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.ListPostsInput
		expectInt  int
		expectBool bool
	}{
		{"true", &types.ListPostsInput{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.ListPostsInput{}, 0, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			i, b := c.p.PerPageValue()
			asst.Equal(c.expectInt, i)
			asst.Equal(c.expectBool, b)
		})
	}
}

func Test_ListPostsInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.ListPostsInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.ListPostsInput{
				TeamName: "test-team",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "with q",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				Q:        "query",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "q", Value: "query"},
				},
			},
		},
		{
			name: "with include",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				Include:  "include1,include2",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "include", Value: "include1,include2"},
				},
			},
		},
		{
			name: "with page",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
				},
			},
		},
		{
			name: "with per_page",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "with sort",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				Sort:     types.ListPostsSortBestMatch,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "sort", Value: "best_match"},
				},
			},
		},
		{
			name: "with order",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				Order:    types.ListPostsOrderAsc,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "order", Value: "asc"},
				},
			},
		},
		{
			name: "with all",
			p: &types.ListPostsInput{
				TeamName: "test-team",
				Q:        "query",
				Include:  "include",
				Sort:     types.ListPostsSortCreated,
				Order:    types.ListPostsOrderDesc,
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "q", Value: "query"},
					{Key: "include", Value: "include"},
					{Key: "sort", Value: "created"},
					{Key: "order", Value: "desc"},
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "ng: not has required parameter",
			p: &types.ListPostsInput{
				Sort:    types.ListPostsSortCreated,
				Order:   types.ListPostsOrderDesc,
				Page:    gesa.NewPageNumber(1),
				PerPage: gesa.NewPageNumber(2),
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: nil",
			p:       nil,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}

func Test_GetPostInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.GetPostInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.GetPostInput{
				TeamName:   "test-team",
				PostNumber: 1,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "with include",
			p: &types.GetPostInput{
				TeamName:   "test-team",
				PostNumber: 1,
				Include:    "include1,include2",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{
					{Key: "include", Value: "include1,include2"},
				},
			},
		},
		{
			name: "ng: not has required parameter: team_name is empty",
			p: &types.GetPostInput{
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.GetPostInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: both are empty",
			p:       &types.GetPostInput{},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: nil",
			p:       nil,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}

func Test_CreatePostInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.CreatePostInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.CreatePostInput{
				TeamName: "test-team",
				Name:     "test-post",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"post":{"name":"test-post"}}`),
			},
		},
		{
			name: "ng: not has required parameter: team_name is empty",
			p: &types.CreatePostInput{
				Name: "test-post",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_name is empty",
			p: &types.CreatePostInput{
				TeamName: "test-team",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: both are empty",
			p:       &types.CreatePostInput{},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: nil",
			p:       nil,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}

func Test_UpdatePostInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.UpdatePostInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.UpdatePostInput{
				TeamName:   "test-team",
				PostNumber: 1,
				BodyMD:     gesa.String("body"),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  strings.NewReader(`{"post":{"body_md":"body"}}`),
			},
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.UpdatePostInput{
				TeamName: "test-post",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_name is empty",
			p: &types.UpdatePostInput{
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: both are empty",
			p: &types.UpdatePostInput{
				BodyMD: gesa.String("body"),
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: nil",
			p:       nil,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}

func Test_DeletePostInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.DeletePostInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p: &types.DeletePostInput{
				TeamName:   "test-team",
				PostNumber: 1,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team_name", Value: "test-team"},
					{Key: ":post_number", Value: "1"},
				},
				Query: internal.QueryParameterList{},
				Body:  nil,
			},
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.DeletePostInput{
				TeamName: "test-post",
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name: "ng: not has required parameter: post_name is empty",
			p: &types.DeletePostInput{
				PostNumber: 1,
			},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: not has required parameter: both are empty",
			p:       &types.DeletePostInput{},
			expect:  nil,
			wantErr: true,
		},
		{
			name:    "ng: nil",
			p:       nil,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep, err := c.p.EsaAPIParameter()
			asst := assert.New(tt)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			asst.Equal(c.expect, ep)
		})
	}
}
