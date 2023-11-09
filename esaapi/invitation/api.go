package invitation

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/invitation/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	getURLInvitationEndpoint        = "https://api.esa.io/:esa_api_version/teams/:team_name/invitation"
	regenerateURLInvitationEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/invitation_regenerator"
	createEmailInvitationsEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/invitations"
	listEmailInvitationsEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/invitations"
	deleteEmailInvitationEndpoint   = "https://api.esa.io/:esa_api_version/teams/:team_name/invitations/:code"
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

// ListEmailInvitations calls getting list of all email invitations API.
// POST /v1/teams/:team_name/invitations
func ListEmailInvitations(ctx context.Context, c *gesa.Client, p *types.ListEmailInvitationsInput) (*types.ListEmailInvitationsOutput, error) {
	res := &types.ListEmailInvitationsOutput{}
	if err := c.CallAPI(ctx, listEmailInvitationsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateEmailInvitations calls create email invitations API.
// POST /v1/teams/:team_name/invitations
func CreateEmailInvitations(ctx context.Context, c *gesa.Client, p *types.CreateEmailInvitationsInput) (*types.CreateEmailInvitationsOutput, error) {
	res := &types.CreateEmailInvitationsOutput{}
	if err := c.CallAPI(ctx, createEmailInvitationsEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteEmailInvitation calls delete email invitation API.
// DELETE /v1/teams/:team_name/invitations/:code
func DeleteEmailInvitation(ctx context.Context, c *gesa.Client, p *types.DeleteEmailInvitationInput) (*types.DeleteEmailInvitationOutput, error) {
	res := &types.DeleteEmailInvitationOutput{}
	if err := c.CallAPI(ctx, deleteEmailInvitationEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
