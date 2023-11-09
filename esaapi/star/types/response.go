package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/gesa"
)

// ListPostStargazersOutput is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number/stargazers
type ListPostStargazersOutput struct {
	Stargazers []models.Stargazer `json:"stargazers"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListPostStargazersOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CreatePostStarOutput is struct for the response of
// POST /v1/teams/:team_name/posts/:post_number/star
type CreatePostStarOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreatePostStarOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// DeletePostStarOutput is struct for the response of
// DELETE /v1/teams/:team_name/posts/:post_number/star
type DeletePostStarOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *DeletePostStarOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// ListCommentStargazersOutput is struct for the response of
// GET /v1/teams/:team_name/comments/:comment_id/stargazers
type ListCommentStargazersOutput struct {
	Stargazers []models.Stargazer `json:"stargazers"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListCommentStargazersOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CreateCommentStarOutput is struct for the response of
// POST /v1/teams/:team_name/comments/:comment_id/star
type CreateCommentStarOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreateCommentStarOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// DeleteCommentStarOutput is struct for the response of
// DELETE /v1/teams/:team_name/comments/:comment_id/star
type DeleteCommentStarOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *DeleteCommentStarOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
