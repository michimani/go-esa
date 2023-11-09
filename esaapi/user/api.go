package user

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/user/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	getMeEndpoint = "https://api.esa.io/:esa_api_version/user"
)

// GetMe calls getting user using API.
// GET /v1/user
func GetMe(ctx context.Context, c *gesa.Client, p *types.GetMeInput) (*types.GetMeOutput, error) {
	res := &types.GetMeOutput{}
	if err := c.CallAPI(ctx, getMeEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
