package types_test

import (
	"io"
	"testing"

	"github.com/michimani/go-esa/esaapi/teams/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/stretchr/testify/assert"
)

func Test_TeamsGetParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.TeamsGetParam
		want    io.Reader
		wantErr bool
	}{
		{"ok, nil", nil, nil, false},
		{"ok: empty", &types.TeamsGetParam{}, nil, false},
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

func Test_TeamsGetParam_PageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.TeamsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.TeamsGetParam{Page: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.TeamsGetParam{}, 0, false},
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

func Test_TeamsGetParam_PerPageValue(t *testing.T) {
	cases := []struct {
		name       string
		p          *types.TeamsGetParam
		expectInt  int
		expectBool bool
	}{
		{"true", &types.TeamsGetParam{PerPage: gesa.NewPageNumber(1)}, 1, true},
		{"false", &types.TeamsGetParam{}, 0, false},
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

func Test_TeamsGetParam_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.TeamsGetParam
		expect string
	}{
		{
			name:   "ok",
			params: &types.TeamsGetParam{},
			expect: endpoint,
		},
		{
			name: "with page",
			params: &types.TeamsGetParam{
				Page: gesa.NewPageNumber(1),
			},
			expect: endpoint + "?page=1",
		},
		{
			name: "with per_page",
			params: &types.TeamsGetParam{
				PerPage: gesa.NewPageNumber(2),
			},
			expect: endpoint + "?per_page=2",
		},
		{
			name: "with page",
			params: &types.TeamsGetParam{
				Page:    gesa.NewPageNumber(1),
				PerPage: gesa.NewPageNumber(2),
			},
			expect: endpoint + "?page=1&per_page=2",
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

func Test_TeamsTeamNameGetParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.TeamsTeamNameGetParam
		want    io.Reader
		wantErr bool
	}{
		{"ok, nil", nil, nil, false},
		{"ok: empty", &types.TeamsTeamNameGetParam{}, nil, false},
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

func Test_TeamsTeamNameGetParam_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.TeamsTeamNameGetParam
		expect string
	}{
		{
			name: "ok",
			params: &types.TeamsTeamNameGetParam{
				TeamName: "test-team",
			},
			expect: endpoint,
		},
		{
			name: "ng: empty value",
			params: &types.TeamsTeamNameGetParam{
				TeamName: "",
			},
			expect: "",
		},
		{
			name:   "ng: empty params",
			params: &types.TeamsTeamNameGetParam{},
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

func Test_TeamsTeamNameGetParam_ParameterMap(t *testing.T) {
	cases := []struct {
		name string
		p    *types.TeamsTeamNameGetParam
		want map[string]string
	}{
		{"ok, nil", nil, nil},
		{"ok: empty", &types.TeamsTeamNameGetParam{}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			m := c.p.ParameterMap()
			asst.Equal(c.want, m)
		})
	}
}
