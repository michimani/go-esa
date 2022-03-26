package invitation

import (
	"context"

	"github.com/michimani/go-esa/esaapi/invitation/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	getURLInvitationEndpoint        = "https://api.esa.io/:esa_api_version/teams/:team_name/invitation"
	regenerateURLInvitationEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/invitation_regenerator"
	createEmailInvitations          = "https://api.esa.io/:esa_api_version/teams/:team_name/invitations"
	listEmailInvitations            = "https://api.esa.io/:esa_api_version/teams/:team_name/invitations"
)

// GetURLInvitation calls getting a invitation URL API.
// GET /v1/teams/:team_name/invitation
func GetURLInvitation(ctx context.Context, c *gesa.Client, p *types.GetURLInvitationInput) (*types.GetURLInvitationOutput, error) {
	res := &types.GetURLInvitationOutput{}
	if err := c.CallAPI(ctx, getURLInvitationEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// RegenerateURLInvitation calls regenerating a invitation URL API.
// POST /v1/teams/:team_name/invitation_regenerator
func RegenerateURLInvitation(ctx context.Context, c *gesa.Client, p *types.RegenerateURLInvitationInput) (*types.RegenerateURLInvitationOutput, error) {
	res := &types.RegenerateURLInvitationOutput{}
	if err := c.CallAPI(ctx, regenerateURLInvitationEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ListEmailInvitations calls create email invitations API.
// POST /v1/teams/:team_name/invitations
func ListEmailInvitations(ctx context.Context, c *gesa.Client, p *types.ListEmailInvitationsInput) (*types.ListEmailInvitationsOutput, error) {
	res := &types.ListEmailInvitationsOutput{}
	if err := c.CallAPI(ctx, listEmailInvitations, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateEmailInvitations calls create email invitations API.
// POST /v1/teams/:team_name/invitations
func CreateEmailInvitations(ctx context.Context, c *gesa.Client, p *types.CreateEmailInvitationsInput) (*types.CreateEmailInvitationsOutput, error) {
	res := &types.CreateEmailInvitationsOutput{}
	if err := c.CallAPI(ctx, createEmailInvitations, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
