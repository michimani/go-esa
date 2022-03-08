package types_test

import (
	"io"
	"testing"

	"github.com/michimani/go-esa/esaapi/stats/types"
	"github.com/stretchr/testify/assert"
)

func Test_StatsGetParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.StatsGetParam
		want    io.Reader
		wantErr bool
	}{
		{"ok, nil", nil, nil, false},
		{"ok: empty", &types.StatsGetParam{}, nil, false},
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

func Test_StatsGetParam_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/:team_name"

	cases := []struct {
		name   string
		params *types.StatsGetParam
		expect string
	}{
		{
			name: "ok",
			params: &types.StatsGetParam{
				TeamName: "test-team",
			},
			expect: "test/endpoint/test-team",
		},
		{
			name: "ng: empty value",
			params: &types.StatsGetParam{
				TeamName: "",
			},
			expect: "",
		},
		{
			name:   "ng: empty params",
			params: &types.StatsGetParam{},
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

func Test_StatsGetParam_ParameterMap(t *testing.T) {
	cases := []struct {
		name string
		p    *types.StatsGetParam
		want map[string]string
	}{
		{"ok, nil", nil, nil},
		{"ok: empty", &types.StatsGetParam{}, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			m := c.p.ParameterMap()
			asst.Equal(c.want, m)
		})
	}
}
