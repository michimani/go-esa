package types

import "github.com/michimani/go-esa/internal"

type GetOAuthTokenInfoInput struct{}

func (p *GetOAuthTokenInfoInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	return &internal.EsaAPIParameter{}, nil
}
