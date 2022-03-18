package comments

import (
	"context"

	"github.com/michimani/go-esa/esaapi/comments/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	CommentsGetEndpoint = "https://api.esa.io/:esa_api_version//teams/:team_name/posts/:post_number/comments"
)

// CommentsGet calls getting members API.
// GET /:esa_api_version//teams/:team_name/posts/:post_number/comments
func CommentsGet(ctx context.Context, c *gesa.GesaClient, p *types.CommentsGetParam) (*types.CommentsGetResponse, error) {
	res := &types.CommentsGetResponse{}
	if err := c.CallAPI(ctx, CommentsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
