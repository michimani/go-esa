package stars

import (
	"context"

	"github.com/michimani/go-esa/esaapi/stars/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listPostStargazersEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/stargazers"
)

// ListPostStargazers calls getting all stargazers in a post API.
// GET /:esa_api_version/teams/:team_name/posts/:post_number/stargazers
func ListPostStargazers(ctx context.Context, c *gesa.GesaClient, p *types.ListPostStargazersInput) (*types.ListPostStargazersOutput, error) {
	res := &types.ListPostStargazersOutput{}
	if err := c.CallAPI(ctx, listPostStargazersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
