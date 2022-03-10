package types_test

import (
	"io"
	"testing"

	"github.com/michimani/go-esa/esaapi/members/types"
	"github.com/michimani/go-esa/gesa"
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

func Test_MembersGetParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.MembersGetParam
		want    io.Reader
		wantErr bool
	}{
		{"ok, nil", nil, nil, false},
		{"ok: empty", &types.MembersGetParam{}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			r, err := c.p.Body()
			if c.wantErr {
				asst.NotNil(err)
				asst.Nil(r)
				return
			}

			asst.Nil(err)
			asst.Equal(c.want, r)
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

func Test_MembersGetParam_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.MembersGetParam
		expect string
	}{
		{
			name: "ok",
			params: &types.MembersGetParam{
				TeamName: "test-team",
			},
			expect: endpoint,
		},
		{
			name: "with page",
			params: &types.MembersGetParam{
				TeamName: "test-team",
				Page:     gesa.NewPageNumber(1),
			},
			expect: endpoint + "?page=1",
		},
		{
			name: "with per_page",
			params: &types.MembersGetParam{
				TeamName: "test-team",
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: endpoint + "?per_page=2",
		},
		{
			name: "with sort",
			params: &types.MembersGetParam{
				TeamName: "test-team",
				Sort:     types.MembersGetSortJoined,
			},
			expect: endpoint + "?sort=joined",
		},
		{
			name: "with order",
			params: &types.MembersGetParam{
				TeamName: "test-team",
				Order:    types.MembersGetOrderAsc,
			},
			expect: endpoint + "?order=asc",
		},
		{
			name: "with all",
			params: &types.MembersGetParam{
				TeamName: "test-team",
				Sort:     types.MembersGetSortPostsCount,
				Order:    types.MembersGetOrderDesc,
				Page:     gesa.NewPageNumber(1),
				PerPage:  gesa.NewPageNumber(2),
			},
			expect: endpoint + "?order=desc&page=1&per_page=2&sort=posts_count",
		},
		{
			name: "ng: not has required parameter",
			params: &types.MembersGetParam{
				Sort:    types.MembersGetSortPostsCount,
				Order:   types.MembersGetOrderDesc,
				Page:    gesa.NewPageNumber(1),
				PerPage: gesa.NewPageNumber(2),
			},
			expect: "",
		},
		{
			name:   "ng: nil",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_MembersScreenNameDeleteParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.MembersScreenNameDeleteParam
		want    io.Reader
		wantErr bool
	}{
		{"ok, nil", nil, nil, false},
		{"ok: empty", &types.MembersScreenNameDeleteParam{}, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			r, err := c.p.Body()
			if c.wantErr {
				asst.NotNil(err)
				asst.Nil(r)
				return
			}

			asst.Nil(err)
			asst.Equal(c.want, r)
		})
	}
}

func Test_MembersScreenNameDeleteParam_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/:team_name/:screen_name"

	cases := []struct {
		name   string
		params *types.MembersScreenNameDeleteParam
		expect string
	}{
		{
			name: "ok",
			params: &types.MembersScreenNameDeleteParam{
				TeamName:   "test-team",
				ScreenName: "test-screen-name",
			},
			expect: "test/endpoint/test-team/test-screen-name",
		},
		{
			name: "ng: empty value: team_name",
			params: &types.MembersScreenNameDeleteParam{
				TeamName:   "",
				ScreenName: "test-screen-name",
			},
			expect: "",
		},
		{
			name: "ng: empty value: screen_name",
			params: &types.MembersScreenNameDeleteParam{
				TeamName:   "test-team",
				ScreenName: "",
			},
			expect: "",
		},
		{
			name: "ng: empty value: both",
			params: &types.MembersScreenNameDeleteParam{
				TeamName:   "",
				ScreenName: "",
			},
			expect: "",
		},
		{
			name:   "ng: empty params",
			params: &types.MembersScreenNameDeleteParam{},
			expect: "",
		},
		{
			name:   "ng: nil",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_MembersScreenNameDeleteParam_ParameterMap(t *testing.T) {
	cases := []struct {
		name string
		p    *types.MembersScreenNameDeleteParam
		want map[string]string
	}{
		{"ok, nil", nil, nil},
		{"ok: empty", &types.MembersScreenNameDeleteParam{}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			m := c.p.ParameterMap()
			asst.Equal(c.want, m)
		})
	}
}
