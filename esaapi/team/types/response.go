package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/gesa"
)

type ListTeamsOutput struct {
	Teams      []models.Team    `json:"teams"`
	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListTeamsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

type GetTeamOutput struct {
	models.Team

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetTeamOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
