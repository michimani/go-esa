package types_test

import (
	"io"
	"testing"

	"github.com/michimani/go-esa/esaapi/oauth/types"
	"github.com/stretchr/testify/assert"
)

func Test_OAuthTokenInfoGetParam_Body(t *testing.T) {
	cases := []struct {
		name    string
		p       *types.OAuthTokenInfoGetParam
		expect  io.Reader
		wantErr bool
	}{
		{
			name:   "ok, nil",
			p:      nil,
			expect: nil,
		},
		{
			name:   "ok: empty",
			p:      &types.OAuthTokenInfoGetParam{},
			expect: nil,
		},
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
			asst.Equal(c.expect, r)
		})
	}
}

func Test_OAuthTokenInfoGetParam_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.OAuthTokenInfoGetParam
		base   string
		expect string
	}{
		{
			name:   "ok, nil",
			p:      nil,
			base:   "endpoint",
			expect: "endpoint",
		},
		{
			name:   "ok: empty",
			p:      &types.OAuthTokenInfoGetParam{},
			base:   "endpoint",
			expect: "endpoint",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			s := c.p.ResolveEndpoint(c.base)
			asst.Equal(c.expect, s)
		})
	}
}
func Test_OAuthTokenInfoGetParam_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.OAuthTokenInfoGetParam
		expect map[string]string
	}{
		{
			name:   "ok, nil",
			p:      nil,
			expect: nil,
		},
		{
			name:   "ok: empty",
			p:      &types.OAuthTokenInfoGetParam{},
			expect: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			m := c.p.ParameterMap()
			asst.Equal(c.expect, m)
		})
	}
}
