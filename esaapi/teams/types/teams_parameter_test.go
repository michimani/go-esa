package types_test

import (
	"testing"

	"github.com/michimani/go-esa/esaapi/teams/types"
	"github.com/michimani/go-esa/gesa"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

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

func Test_TeamsGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.TeamsGetParam
		expect *internal.EsaAPIParameter
	}{
		{
			name: "ok",
			p:    &types.TeamsGetParam{},
			expect: &internal.EsaAPIParameter{
				Path:  internal.PathParameterList{},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "with page",
			p: &types.TeamsGetParam{
				Page: gesa.NewPageNumber(1),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
				},
			},
		},
		{
			name: "with per_page",
			p: &types.TeamsGetParam{
				PerPage: gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{},
				Query: internal.QueryParameterList{
					{Key: "per_page", Value: "2"},
				},
			},
		},
		{
			name: "with all",
			p: &types.TeamsGetParam{
				Page:    gesa.NewPageNumber(1),
				PerPage: gesa.NewPageNumber(2),
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{},
				Query: internal.QueryParameterList{
					{Key: "page", Value: "1"},
					{Key: "per_page", Value: "2"},
				},
			},
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

func Test_TeamsTeamNameGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.TeamsTeamNameGetParam
		expect *internal.EsaAPIParameter
	}{
		{
			name: "ok",
			p: &types.TeamsTeamNameGetParam{
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
			name:   "ng: not has required parameter",
			p:      &types.TeamsTeamNameGetParam{},
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
