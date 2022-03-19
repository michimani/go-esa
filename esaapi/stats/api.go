package stats

import (
	"context"

	"github.com/michimani/go-esa/esaapi/stats/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	getStatsEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/stats"
)

// GetStats calls getting stats API.
// GET /:esa_api_version/teams/:team_name/stats
func GetStats(ctx context.Context, c *gesa.GesaClient, p *types.GetStatsInput) (*types.GetStatsOutput, error) {
	res := &types.GetStatsOutput{}
	if err := c.CallAPI(ctx, getStatsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
