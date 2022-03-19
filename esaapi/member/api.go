package member

import (
	"context"

	"github.com/michimani/go-esa/esaapi/member/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listMembersEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/members"
	deleteMemberEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/members/:screen_name"
)

// ListMembers calls getting members API.
// GET /:esa_api_version/teams/:team_name/members
func ListMembers(ctx context.Context, c *gesa.GesaClient, p *types.ListMembersInput) (*types.ListMembersOutput, error) {
	res := &types.ListMembersOutput{}
	if err := c.CallAPI(ctx, listMembersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteMember calls deleting a member API.
// DELETE /:esa_api_version/teams/:team_name/members/:screen_name
func DeleteMember(ctx context.Context, c *gesa.GesaClient, p *types.DeleteMemberInput) (*types.DeleteMemberOutput, error) {
	res := &types.DeleteMemberOutput{}
	if err := c.CallAPI(ctx, deleteMemberEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
