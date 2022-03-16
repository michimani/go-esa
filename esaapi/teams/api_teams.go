package teams

import (
	"context"

	"github.com/michimani/go-esa/esaapi/teams/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	TeamsGetEndpoint         = "https://api.esa.io/:esa_api_version/teams"
	TeamsTeamNameGetEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name"
)

// TeamsGet calls getting teams API.
// GET /:esa_api_version/teams
func TeamsGet(ctx context.Context, c *gesa.GesaClient, p *types.TeamsGetParam) (*types.TeamsGetResponse, error) {
	res := &types.TeamsGetResponse{}
	if err := c.CallAPI(ctx, TeamsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// TeamsTeamNameGet calls getting a team API.
// GET /:esa_api_version/teams/:team_name
func TeamsTeamNameGet(ctx context.Context, c *gesa.GesaClient, p *types.TeamsTeamNameGetParam) (*types.TeamsTeamNameGetResponse, error) {
	res := &types.TeamsTeamNameGetResponse{}
	if err := c.CallAPI(ctx, TeamsTeamNameGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
