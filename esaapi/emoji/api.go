package emoji

import (
	"context"

	"github.com/michimani/go-esa/v2/esaapi/emoji/types"
	"github.com/michimani/go-esa/v2/gesa"
)

const (
	listEmojisEndpoint  = "https://api.esa.io/:esa_api_version/teams/:team_name/emojis"
	createEmojiEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/emojis"
	deleteEmojiEndpoint = "https://api.esa.io/:esa_api_version/teams/:team_name/emojis/:code"
)

// ListEmojis calls getting all emojis in the team API.
// GET /v1/teams/:team_name/emojis
func ListEmojis(ctx context.Context, c *gesa.Client, p *types.ListEmojisInput) (*types.ListEmojisOutput, error) {
	res := &types.ListEmojisOutput{}
	if err := c.CallAPI(ctx, listEmojisEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateEmoji calls creating new emoji API.
// POST /v1/teams/:team_name/emojis
func CreateEmoji(ctx context.Context, c *gesa.Client, p *types.CreateEmojiInput) (*types.CreateEmojiOutput, error) {
	res := &types.CreateEmojiOutput{}
	if err := c.CallAPI(ctx, createEmojiEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteEmoji calls deleting a emoji API.
// DELETE /v1/teams/:team_name/emojis/:code
func DeleteEmoji(ctx context.Context, c *gesa.Client, p *types.DeleteEmojiInput) (*types.DeleteEmojiOutput, error) {
	res := &types.DeleteEmojiOutput{}
	if err := c.CallAPI(ctx, deleteEmojiEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
