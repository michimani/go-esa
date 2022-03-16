package oauth

import (
	"context"

	"github.com/michimani/go-esa/esaapi/oauth/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	OAuthTokenInfoEndpoint = "https://api.esa.io/oauth/token/info"
)

// OAuthTokenInfoGet calls getting OAuth token information API.
// GET /oauth/token/info
func OAuthTokenInfoGet(ctx context.Context, c *gesa.GesaClient, p *types.OAuthTokenInfoGetParam) (*types.OAuthTokenInfoGetResponse, error) {
	res := &types.OAuthTokenInfoGetResponse{}
	if err := c.CallAPI(ctx, OAuthTokenInfoEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
