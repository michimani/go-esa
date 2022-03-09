package types

import (
	"net/http"
	"time"

	"github.com/michimani/go-esa/gesa"
)

type MembersGetResponse struct {
	Members []Member `json:"members"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

type Member struct {
	Myself         bool       `json:"myself"`
	Name           string     `json:"name"`
	ScreenName     string     `json:"screen_name"`
	Icon           string     `json:"icon"`
	Role           string     `json:"role"`
	PostsCount     int        `json:"posts_count"`
	JoinedAt       *time.Time `json:"joined_at"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
	Email          string     `json:"email,omitempty"`
}

func (r *MembersGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
