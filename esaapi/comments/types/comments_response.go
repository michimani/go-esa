package types

import (
	"net/http"

	"github.com/michimani/go-esa/esaapi/models"
	"github.com/michimani/go-esa/gesa"
)

// CommentsGetResponse is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number/comments
type CommentsGetResponse struct {
	Comments []models.Comment `json:"comments"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CommentsGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CommentsCommentIDGetResponse is struct for the response of
// GET /v1/teams/:team_name/comments/:comment_id
type CommentsCommentIDGetResponse struct {
	models.Comment

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CommentsCommentIDGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CommentsPostResponse is struct for the response of
// POST /v1/teams/:team_name/posts/:post_number/comments
type CommentsPostResponse struct {
	models.Comment

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CommentsPostResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CommentsCommentIDPatchResponse is struct for the response of
// PATCH /v1/teams/:team_name/comments/:comment_id
type CommentsCommentIDPatchResponse struct {
	models.Comment

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CommentsCommentIDPatchResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CommentsCommentIDDeleteResponse is struct for the response of
// DELETE /v1/teams/:team_name/comments/:comment_id
type CommentsCommentIDDeleteResponse struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CommentsCommentIDDeleteResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
