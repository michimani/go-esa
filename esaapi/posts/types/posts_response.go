package types

import (
	"net/http"

	"github.com/michimani/go-esa/esaapi/models"
	"github.com/michimani/go-esa/gesa"
)

// PostsGetResponse is struct for the response of
// GET /v1/teams/:team_name/posts
type PostsGetResponse struct {
	Posts []models.Post `json:"posts"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *PostsGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// PostsPostNumberGetResponse is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number
type PostsPostNumberGetResponse struct {
	models.Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *PostsPostNumberGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// PostsPostResponse is struct for the response of
// POST /v1/teams/:team_name/posts
type PostsPostResponse struct {
	models.Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *PostsPostResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// PostsPostNumberPatchResponse is struct for the response of
// PATCH /v1/teams/:team_name/posts/:post_number
type PostsPostNumberPatchResponse struct {
	models.Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *PostsPostNumberPatchResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
