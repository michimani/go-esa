package watch

import (
	"context"

	"github.com/michimani/go-esa/esaapi/watch/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listWatchersEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/watchers"
	createWatchEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/watch"
)

// ListWatchers calls getting all watchers in a post API.
// GET /v1/teams/:team_name/posts/:post_number/watchers
func ListWatchers(ctx context.Context, c *gesa.Client, p *types.ListWatchersInput) (*types.ListWatchersOutput, error) {
	res := &types.ListWatchersOutput{}
	if err := c.CallAPI(ctx, listWatchersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateWatch calls create watch for a post API.
// POST /v1/teams/:team_name/posts/:post_number/watch
func CreateWatch(ctx context.Context, c *gesa.Client, p *types.CreateWatchInput) (*types.CreateWatchOutput, error) {
	res := &types.CreateWatchOutput{}
	if err := c.CallAPI(ctx, createWatchEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
