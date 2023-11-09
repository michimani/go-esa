package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/gesa"
)

// ListPostCommentsOutput is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number/comments
type ListPostCommentsOutput struct {
	Comments []models.Comment `json:"comments"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListPostCommentsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// GetCommentOutput is struct for the response of
// GET /v1/teams/:team_name/comments/:comment_id
type GetCommentOutput struct {
	models.Comment

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetCommentOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CreateCommentOutput is struct for the response of
// POST /v1/teams/:team_name/posts/:post_number/comments
type CreateCommentOutput struct {
	models.Comment

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreateCommentOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// UpdateCommentOutput is struct for the response of
// PATCH /v1/teams/:team_name/comments/:comment_id
type UpdateCommentOutput struct {
	models.Comment

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *UpdateCommentOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// DeleteCommentOutput is struct for the response of
// DELETE /v1/teams/:team_name/comments/:comment_id
type DeleteCommentOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *DeleteCommentOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// ListTeamCommentsOutput is struct for the response of
// GET /v1/teams/:team_name/comments
type ListTeamCommentsOutput struct {
	Comments []models.Comment `json:"comments"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListTeamCommentsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
