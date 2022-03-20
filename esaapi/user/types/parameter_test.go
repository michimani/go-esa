package types_test

import (
	"testing"

	"github.com/michimani/go-esa/esaapi/user/types"
	"github.com/michimani/go-esa/internal"
	"github.com/stretchr/testify/assert"
)

func Test_GetMeInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.GetMeInput
		expect  *internal.EsaAPIParameter
		wantErr bool
	}{
		{
			name: "ok",
			p:    &types.GetMeInput{},
			expect: &internal.EsaAPIParameter{
				Path:  internal.PathParameterList{},
				Query: internal.QueryParameterList{},
			},
		},
		{
			name: "with include",
			p: &types.GetMeInput{
				Include: "test-include",
			},
			expect: &internal.EsaAPIParameter{
				Path: internal.PathParameterList{},
				Query: internal.QueryParameterList{
					{Key: "include", Value: "test-include"},
				},
			},
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
			asst := assert.New(tt)
			ep, err := c.p.EsaAPIParameter()
			if c.wantErr {
				asst.Error(err)
				asst.Nil(ep)
				return
			}
			assert.Equal(tt, c.expect, ep)
		})
	}
}
