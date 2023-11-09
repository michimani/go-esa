package post

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/post/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	listPostsEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/posts"
	getPostEndpoint    = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number"
	createPostEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts"
	updatePostEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number"
	deletePostEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/posts/:post_number"
)

// ListPosts calls getting posts API.
// GET /:esa_api_version/teams/:team_name/posts
func ListPosts(ctx context.Context, c *gesa.Client, p *types.ListPostsInput) (*types.ListPostsOutput, error) {
	res := &types.ListPostsOutput{}
	if err := c.CallAPI(ctx, listPostsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPost calls getting a post API.
// GET /:esa_api_version/teams/:team_name/posts/:post_number
func GetPost(ctx context.Context, c *gesa.Client, p *types.GetPostInput) (*types.GetPostOutput, error) {
	res := &types.GetPostOutput{}
	if err := c.CallAPI(ctx, getPostEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreatePost calls creating a new post API.
// POST /:esa_api_version/teams/:team_name/posts
func CreatePost(ctx context.Context, c *gesa.Client, p *types.CreatePostInput) (*types.CreatePostOutput, error) {
	res := &types.CreatePostOutput{}
	if err := c.CallAPI(ctx, createPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdatePost calls updating a post API.
// PATCH /:esa_api_version/teams/:team_name/posts/:post_number
func UpdatePost(ctx context.Context, c *gesa.Client, p *types.UpdatePostInput) (*types.UpdatePostOutput, error) {
	res := &types.UpdatePostOutput{}
	if err := c.CallAPI(ctx, updatePostEndpoint, "PATCH", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeletePost calls updating a post API.
// DELETE /:esa_api_version/teams/:team_name/posts/:post_number
func DeletePost(ctx context.Context, c *gesa.Client, p *types.DeletePostInput) (*types.DeletePostOutput, error) {
	res := &types.DeletePostOutput{}
	if err := c.CallAPI(ctx, deletePostEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
