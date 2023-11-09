package member

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/member/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	listMembersEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/members"
	deleteMemberEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/members/:screen_name_or_email"
)

// ListMembers calls getting members API.
// GET /:esa_api_version/teams/:team_name/members
func ListMembers(ctx context.Context, c *gesa.Client, p *types.ListMembersInput) (*types.ListMembersOutput, error) {
	res := &types.ListMembersOutput{}
	if err := c.CallAPI(ctx, listMembersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteMember calls deleting a member API.
// DELETE /:esa_api_version/teams/:team_name/members/:screen_name_or_email
func DeleteMember(ctx context.Context, c *gesa.Client, p *types.DeleteMemberInput) (*types.DeleteMemberOutput, error) {
	res := &types.DeleteMemberOutput{}
	if err := c.CallAPI(ctx, deleteMemberEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
