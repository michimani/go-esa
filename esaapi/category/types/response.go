package types

import (
	"net/http"

	"github.com/michimani/go-esa/v2/gesa"
)

// BatchMoveOutput is struct for the response of
// POST v1/teams/:team_name/categories/batch_move
type BatchMoveOutput struct {
	Count int    `json:"count"`
	From  string `json:"from"`
	To    string `json:"to"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *BatchMoveOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
