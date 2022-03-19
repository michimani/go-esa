package watches

import (
	"context"

	"github.com/michimani/go-esa/esaapi/watches/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listWatchersEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/watchers"
)

// ListWatchers calls getting all watchers in a post API.
// GET /v1/teams/:team_name/posts/:post_number/watchers
func ListWatchers(ctx context.Context, c *gesa.GesaClient, p *types.ListWatchersInput) (*types.ListWatchersOutput, error) {
	res := &types.ListWatchersOutput{}
	if err := c.CallAPI(ctx, listWatchersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
