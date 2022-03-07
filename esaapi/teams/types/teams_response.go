package types

import (
	"net/http"

	"github.com/michimani/go-esa/gesa"
)

type TeamsGetResponse struct {
	Teams      []Team           `json:"teams"`
	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

type Team struct {
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
}

func (r *TeamsGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
