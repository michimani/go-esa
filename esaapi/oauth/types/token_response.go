package types

import (
	"net/http"

	"github.com/michimani/go-esa/gesa"
)

type OAuthTokenInfoGetResponse struct {
	Status        string                     `json:"-"`
	StatusCode    int                        `json:"-"`
	RateLimitInfo *gesa.RateLimitInformation `json:"-"`

	ResourceOwnerID *int            `json:"resource_owner_id,omitempty"`
	Scope           []string        `json:"scope,omitempty"`
	ExpiresIn       *gesa.Timestamp `json:"expires_in,omitempty"`
	Application     *Application    `json:"application,omitempty"`
	CreatedAt       *gesa.Timestamp `json:"created_at,omitempty"`
	User            *User           `json:"user,omitempty"`
	Error           *string         `json:"error,omitempty"`
	Message         *string         `json:"message,omitempty"`
}

type Application struct {
	UID string `json:"uid"`
}

type User struct {
	ID int `json:"id"`
}

func (r *OAuthTokenInfoGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
