package types

import "github.com/michimani/go-esa/internal"

type OAuthTokenInfoGetParam struct{}

func (p *OAuthTokenInfoGetParam) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	return &internal.EsaAPIParameter{}, nil
}
