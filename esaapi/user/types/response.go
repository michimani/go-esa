package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/gesa"
)

type GetMeOutput struct {
	models.Me

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetMeOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
