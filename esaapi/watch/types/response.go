package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/gesa"
)

// ListWatchersOutput is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number/watchers
type ListWatchersOutput struct {
	Watchers []models.Watcher `json:"watchers"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListWatchersOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// CreateWatchOutput is struct for the response of
// POST /v1/teams/:team_name/posts/:post_number/watch
type CreateWatchOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreateWatchOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// DeleteWatchOutput is struct for the response of
// DELETE /v1/teams/:team_name/posts/:post_number/watch
type DeleteWatchOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *DeleteWatchOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
