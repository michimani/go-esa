package types_test

import (
	"testing"

	"github.com/michimani/go-esa/v2/esaapi/oauthtoken/types"
	"github.com/michimani/go-esa/v2/internal"
	"github.com/stretchr/testify/assert"
)

func Test_GetOAuthTokenInfoInput_EsaAPIParameter(t *testing.T) {
	cases := []struct {
		name   string
		p      *types.GetOAuthTokenInfoInput
		expect *internal.EsaAPIParameter
	}{
		{
			name:   "ok, nil",
			p:      nil,
			expect: &internal.EsaAPIParameter{},
		},
		{
			name:   "ok: empty",
			p:      &types.GetOAuthTokenInfoInput{},
			expect: &internal.EsaAPIParameter{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			s, err := c.p.EsaAPIParameter()
			asst.NoError(err)
			asst.Equal(c.expect, s)
		})
	}
}
