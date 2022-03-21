package types

import (
	"net/http"

	"github.com/michimani/go-esa/esaapi/models"
	"github.com/michimani/go-esa/gesa"
)

type ListEmojisOutput struct {
	Emojis []models.Emoji `json:"emojis"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *ListEmojisOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
