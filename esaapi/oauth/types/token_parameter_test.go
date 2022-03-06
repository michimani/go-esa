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
		want    io.Reader
		wantErr bool
	}{
		{
			name: "ok, nil",
			p:    nil,
			want: nil,
		},
		{
			name: "ok: empty",
			p:    &types.OAuthTokenInfoGetParam{},
			want: nil,
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
			asst.Equal(c.want, r)
		})
	}
}
