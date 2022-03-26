package types

import (
	"net/http"

	"github.com/michimani/go-esa/gesa"
)

type GetInvitationOutput struct {
	URL string `json:"url"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetInvitationOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
