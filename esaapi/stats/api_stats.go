package stats

import (
	"context"

	"github.com/michimani/go-esa/esaapi/stats/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	StatsGetEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/stats"
)

func StatsGet(ctx context.Context, c *gesa.GesaClient, p *types.StatsGetParam) (*types.StatsGetResponse, error) {
	res := &types.StatsGetResponse{}
	if err := c.CallAPI(ctx, StatsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
