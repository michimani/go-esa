package stars

import (
	"context"

	"github.com/michimani/go-esa/esaapi/stars/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	postNumberStargazersGetEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/stargazers"
)

// PostNumberStargazersGet calls getting all stargazers in a post API.
// GET /:esa_api_version/teams/:team_name/posts/:post_number/stargazers
func PostNumberStargazersGet(ctx context.Context, c *gesa.GesaClient, p *types.PostNumberStargazersGetParam) (*types.PostNumberStargazersGetResponse, error) {
	res := &types.PostNumberStargazersGetResponse{}
	if err := c.CallAPI(ctx, postNumberStargazersGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
