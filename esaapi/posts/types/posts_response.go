package types

import (
	"net/http"
	"time"

	"github.com/michimani/go-esa/gesa"
)

// PostsGetResponse is struct for the response of
// GET /v1/teams/:team_name/posts
type PostsGetResponse struct {
	Posts []Post `json:"posts"`

	PrevPage   *gesa.PageNumber `json:"prev_page,omitempty"`
	NextPage   *gesa.PageNumber `json:"next_page,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	MaxPerPage int              `json:"max_per_page"`

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

type Post struct {
	Number          int         `json:"number"`
	Name            string      `json:"name"`
	FullName        string      `json:"full_name"`
	Wip             bool        `json:"wip"`
	BodyMD          string      `json:"body_md"`
	BodyHTML        string      `json:"body_html"`
	CreatedAt       *time.Time  `json:"created_at"`
	Message         string      `json:"message"`
	Kind            string      `json:"kind"`
	CommentCount    int         `json:"comment_count"`
	DoneTasksCount  int         `json:"done_tasks_count"`
	URL             string      `json:"url"`
	UpdatedAt       *time.Time  `json:"updated_at"`
	Tags            []string    `json:"tags"`
	Category        string      `json:"category"`
	RevisionNumber  int         `json:"revision_number"`
	CreatedBy       User        `json:"created_by"`
	UpdatedBy       User        `json:"updated_by"`
	StargazersCount int         `json:"stargazers_count"`
	WatchersCount   int         `json:"watchers_count"`
	Star            bool        `json:"star"`
	Watch           bool        `json:"watch"`
	SharingURL      string      `json:"sharing_url"`
	Comments        []Comment   `json:"comments,omitempty"`
	Stargazers      []Stargazer `json:"stargazers,omitempty"`
}

type User struct {
	Myself     bool   `json:"myself"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type Comment struct {
	ID              int         `json:"id"`
	BodyMD          string      `json:"body_md"`
	BodyHTML        string      `json:"body_html"`
	CreatedAt       *time.Time  `json:"created_at"`
	UpdatedAt       *time.Time  `json:"updated_at"`
	PostNumber      int         `json:"post_number"`
	URL             string      `json:"url"`
	CreatedBy       User        `json:"created_by"`
	StargazersCount int         `json:"stargazers_count"`
	Star            bool        `json:"star"`
	Stargazers      []Stargazer `json:"stargazers,omitempty"`
}

type Stargazer struct {
	CreatedAt *time.Time `json:"created_at"`
	Body      string
	User      User
}

func (r *PostsGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}

// PostsPostNumberGetResponse is struct for the response of
// GET /v1/teams/:team_name/posts/:post_number
type PostsPostNumberGetResponse struct {
	Post

	RateLimitInfo *gesa.RateLimitInformation `json:"-"`
}

func (r *PostsPostNumberGetResponse) SetRateLimitInfo(h http.Header) {
	if rri, err := gesa.GetRateLimitInformation(h); err == nil {
		r.RateLimitInfo = rri
	}
}
