package teams

import (
	"context"

	"github.com/michimani/go-esa/esaapi/teams/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	TeamsGetEndpoint = "https://api.esa.io/v1/teams"
)

func TeamsGet(ctx context.Context, c *gesa.GesaClient, p *types.TeamsGetParam) (*types.TeamsGetResponse, error) {
	res := &types.TeamsGetResponse{}
	if err := c.CallAPI(ctx, TeamsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
