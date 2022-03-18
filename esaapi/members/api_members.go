package members

import (
	"context"

	"github.com/michimani/go-esa/esaapi/members/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	membersGetEndpoint              = "https://api.esa.io/:esa_api_version/teams/:team_name/members"
	membersScreenNameDeleteEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/members/:screen_name"
)

// MembersGet calls getting members API.
// GET /:esa_api_version/teams/:team_name/members
func MembersGet(ctx context.Context, c *gesa.GesaClient, p *types.MembersGetParam) (*types.MembersGetResponse, error) {
	res := &types.MembersGetResponse{}
	if err := c.CallAPI(ctx, membersGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// MembersScreenNameDelete calls deleting a member API.
// DELETE /:esa_api_version/teams/:team_name/members/:screen_name
func MembersScreenNameDelete(ctx context.Context, c *gesa.GesaClient, p *types.MembersScreenNameDeleteParam) (*types.MembersScreenNameDeleteResponse, error) {
	res := &types.MembersScreenNameDeleteResponse{}
	if err := c.CallAPI(ctx, membersScreenNameDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
