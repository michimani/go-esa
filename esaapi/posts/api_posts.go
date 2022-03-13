package posts

import (
	"context"

	"github.com/michimani/go-esa/esaapi/posts/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	PostsGetEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts"
)

func PostsGet(ctx context.Context, c *gesa.GesaClient, p *types.PostsGetParam) (*types.PostsGetResponse, error) {
	res := &types.PostsGetResponse{}
	if err := c.CallAPI(ctx, PostsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
