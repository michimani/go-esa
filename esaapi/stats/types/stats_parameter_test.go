package types_test

import (
	"testing"

	"github.com/michimani/go-esa/esaapi/stats/types"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_StatsGetParam_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.StatsGetParam
		expect *internal.EsaAPIParameter
	}{
		{
			name: "ok",
			p: &types.StatsGetParam{
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
			name:   "ng: not has required parameter",
			p:      &types.StatsGetParam{},
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
