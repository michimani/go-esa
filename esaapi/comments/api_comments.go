package comments

import (
	"context"

	"github.com/michimani/go-esa/esaapi/comments/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	listPostCommentsEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/comments"
	getCommentEndpoint       = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id"
	createCommentEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/comments"
	updateCommentEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id"
	deleteCommentEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id"
	listTeamCommentsEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/comments"
)

// ListPostComments calls getting all comments in a post API.
// GET /:esa_api_version/teams/:team_name/posts/:post_number/comments
func ListPostComments(ctx context.Context, c *gesa.GesaClient, p *types.ListPostCommentsInput) (*types.ListPostCommentsOutput, error) {
	res := &types.ListPostCommentsOutput{}
	if err := c.CallAPI(ctx, listPostCommentsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetComment calls getting a comment API.
// GET /:esa_api_version/teams/:team_name/comments/:comment_id
func GetComment(ctx context.Context, c *gesa.GesaClient, p *types.GetCommentInput) (*types.GetCommentOutput, error) {
	res := &types.GetCommentOutput{}
	if err := c.CallAPI(ctx, getCommentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateComment calls creating a comment API.
// POST /:esa_api_version/teams/:team_name/posts/:post_number/comments
func CreateComment(ctx context.Context, c *gesa.GesaClient, p *types.CreateCommentInput) (*types.CreateCommentOutput, error) {
	res := &types.CreateCommentOutput{}
	if err := c.CallAPI(ctx, createCommentEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateComment calls updating a comment API.
// PATCH /:esa_api_version/teams/:team_name/comments/:comment_id
func UpdateComment(ctx context.Context, c *gesa.GesaClient, p *types.UpdateCommentInput) (*types.UpdateCommentOutput, error) {
	res := &types.UpdateCommentOutput{}
	if err := c.CallAPI(ctx, updateCommentEndpoint, "PATCH", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteComment calls deleting a comment API.
// DELETE /:esa_api_version/teams/:team_name/comments/:comment_id
func DeleteComment(ctx context.Context, c *gesa.GesaClient, p *types.DeleteCommentInput) (*types.DeleteCommentOutput, error) {
	res := &types.DeleteCommentOutput{}
	if err := c.CallAPI(ctx, deleteCommentEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ListTeamComments calls getting all comments in team API.
// GET /v1/teams/:team_name/comments
func ListTeamComments(ctx context.Context, c *gesa.GesaClient, p *types.ListTeamCommentsInput) (*types.ListTeamCommentsOutput, error) {
	res := &types.ListTeamCommentsOutput{}
	if err := c.CallAPI(ctx, listTeamCommentsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
