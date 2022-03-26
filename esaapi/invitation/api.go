package invitation

import (
	"context"

	"github.com/michimani/go-esa/esaapi/invitation/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	getInvitationEndpoint        = "https://api.esa.io/:esa_api_version/teams/:team_name/invitation"
	regenerateInvitationEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/invitation_regenerator"
)

// GetInvitation calls getting a invitation API.
// GET /v1/teams/:team_name/invitation
func GetInvitation(ctx context.Context, c *gesa.Client, p *types.GetInvitationInput) (*types.GetInvitationOutput, error) {
	res := &types.GetInvitationOutput{}
	if err := c.CallAPI(ctx, getInvitationEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// RegenerateInvitation calls regenerating a invitation API.
// POST /v1/teams/:team_name/invitation_regenerator
func RegenerateInvitation(ctx context.Context, c *gesa.Client, p *types.RegenerateInvitationInput) (*types.RegenerateInvitationOutput, error) {
	res := &types.RegenerateInvitationOutput{}
	if err := c.CallAPI(ctx, regenerateInvitationEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
