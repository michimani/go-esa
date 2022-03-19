package types

import (
	"net/http"

	"github.com/michimani/go-esa/gesa"
)

type GetStatsOutput struct {
	Members            int `json:"members"`
	Posts              int `json:"posts"`
	PostsWip           int `json:"posts_wip"`
	PostsShipped       int `json:"posts_shipped"`
	Comments           int `json:"comments"`
	Stars              int `json:"stars"`
	DailyActiveUsers   int `json:"daily_active_users"`
	WeeklyActiveUsers  int `json:"weekly_active_users"`
	MonthlyActiveUsers int `json:"monthly_active_users"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *GetStatsOutput) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
