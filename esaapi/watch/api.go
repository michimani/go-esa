package watch

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/watch/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	listWatchersEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/watchers"
	createWatchEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/watch"
	deleteWatchEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/watch"
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

// CreateWatch calls creating watch for a post API.
// POST /v1/teams/:team_name/posts/:post_number/watch
func CreateWatch(ctx context.Context, c *gesa.Client, p *types.CreateWatchInput) (*types.CreateWatchOutput, error) {
	res := &types.CreateWatchOutput{}
	if err := c.CallAPI(ctx, createWatchEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteWatch calls deleting watch for a post API.
// POST /v1/teams/:team_name/posts/:post_number/watch
func DeleteWatch(ctx context.Context, c *gesa.Client, p *types.DeleteWatchInput) (*types.DeleteWatchOutput, error) {
	res := &types.DeleteWatchOutput{}
	if err := c.CallAPI(ctx, deleteWatchEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
