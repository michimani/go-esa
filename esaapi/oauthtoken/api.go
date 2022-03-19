package oauthtoken

import (
	"context"

	"github.com/michimani/go-esa/esaapi/oauthtoken/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	getOAuthTokenInfoEndpoint = "https://api.esa.io/oauth/token/info"
)

// GetOAuthTokenInfo calls getting OAuth token information API.
// GET /oauth/token/info
func GetOAuthTokenInfo(ctx context.Context, c *gesa.Client, p *types.GetOAuthTokenInfoInput) (*types.GetOAuthTokenInfoOutput, error) {
	res := &types.GetOAuthTokenInfoOutput{}
	if err := c.CallAPI(ctx, getOAuthTokenInfoEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
