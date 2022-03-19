package team

import (
	"context"

	"github.com/michimani/go-esa/esaapi/team/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listTeamsEndpoint = "https://api.esa.io/:esa_api_version/teams"
	getTeamEndpoint   = "https://api.esa.io/:esa_api_version/teams/:team_name"
)

// ListTeams calls getting teams API.
// GET /:esa_api_version/teams
func ListTeams(ctx context.Context, c *gesa.Client, p *types.ListTeamsInput) (*types.ListTeamsOutput, error) {
	res := &types.ListTeamsOutput{}
	if err := c.CallAPI(ctx, listTeamsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetTeam calls getting a team API.
// GET /:esa_api_version/teams/:team_name
func GetTeam(ctx context.Context, c *gesa.Client, p *types.GetTeamInput) (*types.GetTeamOutput, error) {
	res := &types.GetTeamOutput{}
	if err := c.CallAPI(ctx, getTeamEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
