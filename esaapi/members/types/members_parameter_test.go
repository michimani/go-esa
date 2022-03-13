package types_test

import (
	"testing"

	"github.com/michimani/go-esa/esaapi/members/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_MembersGetSort_IsValid(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect bool
	}{
		{"ok", "joined", true},
		{"ok", "posts_count", true},
		{"ok", "last_accessed", true},
		{"ng", "unknown", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			asst.Equal(c.expect, types.MembersGetSort(c.s).IsValid())
		})
	}
}

func Test_MembersGetOrder_IsValid(t *testing.T) {
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
			asst.Equal(c.expect, types.MembersGetOrder(c.s).IsValid())
		})
	}
}

func Test_MembersGetParam_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.MembersGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.MembersGetParam{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.MembersGetParam{}, 0, false},
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

func Test_MembersGetParam_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.MembersGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.MembersGetParam{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.MembersGetParam{}, 0, false},
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

func Test_MembersGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.MembersGetParam
		expect *internal.EsaAPIParameter
	}{
		{
			name: "ok",
			p: &types.MembersGetParam{
				TeamName: "test-team",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team", Value: "test-team"},
				},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "with page",
			p: &types.MembersGetParam{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
				},
			},
		},
		{
			name: "with per_page",
			p: &types.MembersGetParam{
				TeamName: "test-team",
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "with sort",
			p: &types.MembersGetParam{
				TeamName: "test-team",
				Sort:     types.MembersGetSortJoined,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "sort", Value: "joined"},
				},
			},
		},
		{
			name: "with order",
			p: &types.MembersGetParam{
				TeamName: "test-team",
				Order:    types.MembersGetOrderAsc,
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "order", Value: "asc"},
				},
			},
		},
		{
			name: "with all",
			p: &types.MembersGetParam{
				TeamName: "test-team",
				Sort:     types.MembersGetSortPostsCount,
				Order:    types.MembersGetOrderDesc,
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{
					{Key: ":team", Value: "test-team"},
				},
				Query: internal.QueryParameterList{
					{Key: "sort", Value: "posts_count"},
					{Key: "order", Value: "desc"},
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "ng: not has required parameter",
			p: &types.MembersGetParam{
				Sort:    types.MembersGetSortPostsCount,
				Order:   types.MembersGetOrderDesc,
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
