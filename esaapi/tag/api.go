package tag

import (
	"context"

	"github.com/michimani/go-esa/esaapi/tag/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listTagsEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/tags"
)

// ListTags calls getting all tags API.
// GET v1/teams/:team_name/tags
func ListTags(ctx context.Context, c *gesa.Client, p *types.ListTagsInput) (*types.ListTagsOutput, error) {
	res := &types.ListTagsOutput{}
	if err := c.CallAPI(ctx, listTagsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
