package posts

import (
	"context"

	"github.com/michimani/go-esa/esaapi/posts/types"
	"github.com/michimani/go-esa/gesa"
)

const (
	PostsGetEndpoint              = "https://api.esa.io/:esa_api_version/teams/:team_name/posts"
	PostsPostNumberGetEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number"
	PostsPostEndpoint             = "https://api.esa.io/:esa_api_version/teams/:team_name/posts"
	PostsPostNumberPatchEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number"
	PostsPostNumberDeleteEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number"
)

// PostsGet calls getting posts API.
// GET /:esa_api_version/teams/:team_name/posts
func PostsGet(ctx context.Context, c *gesa.GesaClient, p *types.PostsGetParam) (*types.PostsGetResponse, error) {
	res := &types.PostsGetResponse{}
	if err := c.CallAPI(ctx, PostsGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PostsPostNumberGet calls getting a post API.
// GET /:esa_api_version/teams/:team_name/posts/:post_number
func PostsPostNumberGet(ctx context.Context, c *gesa.GesaClient, p *types.PostsPostNumberGetParam) (*types.PostsPostNumberGetResponse, error) {
	res := &types.PostsPostNumberGetResponse{}
	if err := c.CallAPI(ctx, PostsPostNumberGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PostsPost calls creating a new post API.
// POST /:esa_api_version/teams/:team_name/posts
func PostsPost(ctx context.Context, c *gesa.GesaClient, p *types.PostsPostParam) (*types.PostsPostResponse, error) {
	res := &types.PostsPostResponse{}
	if err := c.CallAPI(ctx, PostsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PostsPostNumberPatch calls updating a post API.
// PATCH /:esa_api_version/teams/:team_name/posts/:post_number
func PostsPostNumberPatch(ctx context.Context, c *gesa.GesaClient, p *types.PostsPostNumberPatchParam) (*types.PostsPostNumberPatchResponse, error) {
	res := &types.PostsPostNumberPatchResponse{}
	if err := c.CallAPI(ctx, PostsPostNumberPatchEndpoint, "PATCH", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PostsPostNumberDelete calls updating a post API.
// DELETE /:esa_api_version/teams/:team_name/posts/:post_number
func PostsPostNumberDelete(ctx context.Context, c *gesa.GesaClient, p *types.PostsPostNumberDeleteParam) (*types.PostsPostNumberDeleteResponse, error) {
	res := &types.PostsPostNumberDeleteResponse{}
	if err := c.CallAPI(ctx, PostsPostNumberDeleteEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
