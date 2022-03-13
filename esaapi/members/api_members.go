package members

import (
	"context"

	"github.com/michimani/go-esa/esaapi/members/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	MembersGetEndpoint              = "https://api.esa.io/:esa_api_version/teams/:team_name/members"
	MembersScreenNameDeleteEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/members/:screen_name"
)

func MembersGet(ctx context.Context, c *gesa.GesaClient, p *types.MembersGetParam) (*types.MembersGetResponse, error) {
	res := &types.MembersGetResponse{}
	if err := c.CallAPI(ctx, MembersGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func MembersScreenNameDelete(ctx context.Context, c *gesa.GesaClient, p *types.MembersScreenNameDeleteParam) (*types.MembersScreenNameDeleteResponse, error) {
	res := &types.MembersScreenNameDeleteResponse{}
	if err := c.CallAPI(ctx, MembersScreenNameDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}