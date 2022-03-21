package emoji

import (
	"context"

	"github.com/michimani/go-esa/esaapi/emoji/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listEmojisEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/emojis"
)

// ListEmojis calls getting all emojis in the team API.
// GET /v1/teams/:team_name/emojis
func ListEmojis(ctx context.Context, c *gesa.Client, p *types.ListEmojisInput) (*types.ListEmojisOutput, error) {
	res := &types.ListEmojisOutput{}
	if err := c.CallAPI(ctx, listEmojisEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
