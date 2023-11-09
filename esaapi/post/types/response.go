package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/gesa"
)

// ListPostsOutput is struct for the response of
// GET /v1/teams/:team_name/posts
type ListPostsOutput struct {
	Posts []models.Post `json:"posts"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListPostsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// GetPostOutput is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number
type GetPostOutput struct {
	models.Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetPostOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CreatePostOutput is struct for the response of
// POST /v1/teams/:team_name/posts
type CreatePostOutput struct {
	models.Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreatePostOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// UpdatePostOutput is struct for the response of
// PATCH /v1/teams/:team_name/posts/:post_number
type UpdatePostOutput struct {
	models.Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *UpdatePostOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// DeletePostOutput is struct for the response of
// DELETE /v1/teams/:team_name/posts/:post_number
type DeletePostOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *DeletePostOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
