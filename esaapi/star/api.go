package star

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/star/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	listPostStargazersEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/stargazers"
	createPostStarEndpoint        = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/star"
	deletePostStarEndpoint        = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/star"
	listCommentStargazersEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id/stargazers"
	createCommentStarEndpoint     = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id/star"
	deleteCommentStarEndpoint     = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id/star"
)

// ListPostStargazers calls getting all stargazers in a post API.
// GET /v1/teams/:team_name/posts/:post_number/stargazers
func ListPostStargazers(ctx context.Context, c *gesa.Client, p *types.ListPostStargazersInput) (*types.ListPostStargazersOutput, error) {
	res := &types.ListPostStargazersOutput{}
	if err := c.CallAPI(ctx, listPostStargazersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreatePostStar calls getting all stargazers in a post API.
// POST /v1/teams/:team_name/posts/:post_number/star
func CreatePostStar(ctx context.Context, c *gesa.Client, p *types.CreatePostStarInput) (*types.CreatePostStarOutput, error) {
	res := &types.CreatePostStarOutput{}
	if err := c.CallAPI(ctx, createPostStarEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeletePostStar calls getting all stargazers in a post API.
// DELETE /v1/teams/:team_name/posts/:post_number/star
func DeletePostStar(ctx context.Context, c *gesa.Client, p *types.DeletePostStarInput) (*types.DeletePostStarOutput, error) {
	res := &types.DeletePostStarOutput{}
	if err := c.CallAPI(ctx, deletePostStarEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ListCommentStargazers calls getting all stargazers in a comment API.
// GET /v1/teams/:team_name/comments/:comment_id/stargazers
func ListCommentStargazers(ctx context.Context, c *gesa.Client, p *types.ListCommentStargazersInput) (*types.ListCommentStargazersOutput, error) {
	res := &types.ListCommentStargazersOutput{}
	if err := c.CallAPI(ctx, listCommentStargazersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateCommentStar calls getting all stargazers in a comment API.
// POST /v1/teams/:team_name/comments/:comment_id/star
func CreateCommentStar(ctx context.Context, c *gesa.Client, p *types.CreateCommentStarInput) (*types.CreateCommentStarOutput, error) {
	res := &types.CreateCommentStarOutput{}
	if err := c.CallAPI(ctx, createCommentStarEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteCommentStar calls getting all stargazers in a comment API.
// DELETE /v1/teams/:team_name/comments/:comment_id/star
func DeleteCommentStar(ctx context.Context, c *gesa.Client, p *types.DeleteCommentStarInput) (*types.DeleteCommentStarOutput, error) {
	res := &types.DeleteCommentStarOutput{}
	if err := c.CallAPI(ctx, deleteCommentStarEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
