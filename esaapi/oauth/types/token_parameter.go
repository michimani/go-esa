package types

import "io"

type OAuthTokenInfoGetParam struct{}

func (p *OAuthTokenInfoGetParam) Body() (io.Reader, error) {
	return nil, nil
}
