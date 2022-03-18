package comments

import (
	"context"

	"github.com/michimani/go-esa/esaapi/comments/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	commentsGetEndpoint          = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/comments"
	commentsCommentIDGetEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/comments/:comment_id"
	commentsPostEndpoint         = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number/comments"
)

// CommentsGet calls getting members API.
// GET /:esa_api_version//teams/:team_name/posts/:post_number/comments
func CommentsGet(ctx context.Context, c *gesa.GesaClient, p *types.CommentsGetParam) (*types.CommentsGetResponse, error) {
	res := &types.CommentsGetResponse{}
	if err := c.CallAPI(ctx, commentsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CommentsCommentIDGet calls getting members API.
// GET /:esa_api_version//teams/:team_name/posts/:post_number/comments
func CommentsCommentIDGet(ctx context.Context, c *gesa.GesaClient, p *types.CommentsCommentIDGetParam) (*types.CommentsCommentIDGetResponse, error) {
	res := &types.CommentsCommentIDGetResponse{}
	if err := c.CallAPI(ctx, commentsCommentIDGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CommentsPost calls getting members API.
// GET /:esa_api_version//teams/:team_name/posts/:post_number/comments
func CommentsPost(ctx context.Context, c *gesa.GesaClient, p *types.CommentsPostParam) (*types.CommentsPostResponse, error) {
	res := &types.CommentsPostResponse{}
	if err := c.CallAPI(ctx, commentsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
