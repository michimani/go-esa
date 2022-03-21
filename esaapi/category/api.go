package category

import (
	"context"

	"github.com/michimani/go-esa/esaapi/category/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	batchMoveEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/categories/batch_move"
)

// BatchMove calls moving posts from a category to other API.
// POST v1/teams/:team_name/categories/batch_move
func BatchMove(ctx context.Context, c *gesa.Client, p *types.BatchMoveInput) (*types.BatchMoveOutput, error) {
	res := &types.BatchMoveOutput{}
	if err := c.CallAPI(ctx, batchMoveEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
