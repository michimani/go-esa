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

type CreateEmojiOutput struct {
	Code string `json:"code"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *CreateEmojiOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

type DeleteEmojiOutput struct {
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *DeleteEmojiOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
