package types_test

import (
	"testing"

	"github.com/michimani/go-esa/esaapi/posts/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_PostsGetSort_IsValid(t *testing.T) {
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
			asst.Equal(c.expect, types.PostsGetSort(c.s).IsValid())
		})
	}
}

func Test_PostsGetOrder_IsValid(t *testing.T) {
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
			asst.Equal(c.expect, types.PostsGetOrder(c.s).IsValid())
		})
	}
}

func Test_PostsGetParam_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.PostsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.PostsGetParam{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.PostsGetParam{}, 0, false},
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

func Test_PostsGetParam_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.PostsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.PostsGetParam{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.PostsGetParam{}, 0, false},
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

func Test_PostsGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.PostsGetParam
		expect *internal.EsaAPIParameter
	}{
		{
			name: "ok",
			p: &types.PostsGetParam{
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
			p: &types.PostsGetParam{
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
			p: &types.PostsGetParam{
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
			p: &types.PostsGetParam{
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
			p: &types.PostsGetParam{
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
			p: &types.PostsGetParam{
				TeamName: "test-team",
				Sort:     types.PostsGetSortBestMatch,
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
			p: &types.PostsGetParam{
				TeamName: "test-team",
				Order:    types.PostsGetOrderAsc,
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
			p: &types.PostsGetParam{
				TeamName: "test-team",
				Q:        "query",
				Include:  "include",
				Sort:     types.PostsGetSortCreated,
				Order:    types.PostsGetOrderDesc,
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
			p: &types.PostsGetParam{
				Sort:    types.PostsGetSortCreated,
				Order:   types.PostsGetOrderDesc,
				Page:    gesa.NewPageNumber(1),
				PerPage: gesa.NewPageNumber(2),
			},
			expect: nil,
		},
		{
			name:   "ng: nil",
			p:      nil,
			expect: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.p.EsaAPIParameter()
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_PostsPostNumberGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.PostsPostNumberGetParam
		expect *internal.EsaAPIParameter
	}{
		{
			name: "ok",
			p: &types.PostsPostNumberGetParam{
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
			p: &types.PostsPostNumberGetParam{
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
			p: &types.PostsPostNumberGetParam{
				PostNumber: 1,
			},
			expect: nil,
		},
		{
			name: "ng: not has required parameter: post_number is empty",
			p: &types.PostsPostNumberGetParam{
				TeamName: "test-team",
			},
			expect: nil,
		},
		{
			name:   "ng: not has required parameter: both are empty",
			p:      &types.PostsPostNumberGetParam{},
			expect: nil,
		},
		{
			name:   "ng: nil",
			p:      nil,
			expect: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.p.EsaAPIParameter()
			assert.Equal(tt, c.expect, ep)
		})
	}
}
