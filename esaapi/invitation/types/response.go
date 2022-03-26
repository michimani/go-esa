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

type CreateEmailInvitationsOutput struct {
	Invitations []models.EmailInvitations `json:"invitations"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreateEmailInvitationsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
