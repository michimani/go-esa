package types

import (
	"net/http"

	"github.com/michimani/go-esa/esaapi/models"
	"github.com/michimani/go-esa/gesa"
)

type GetURLInvitationOutput struct {
	URL string `json:"url"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetURLInvitationOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

type RegenerateURLInvitationOutput struct {
	URL string `json:"url"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *RegenerateURLInvitationOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

type ListEmailInvitationsOutput struct {
	Invitations []models.EmailInvitations `json:"invitations"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListEmailInvitationsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

type CreateEmailInvitationsOutput struct {
	Invitations []models.EmailInvitations `json:"invitations"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreateEmailInvitationsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
