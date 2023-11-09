package types

import "github.com/michimani/go-esa/v2/internal"

type GetOAuthTokenInfoInput struct{}

func (p *GetOAuthTokenInfoInput) EsaAPIParameter() (*internal.EsaAPIParameter, error) {
	return &internal.EsaAPIParameter{}, nil
}
