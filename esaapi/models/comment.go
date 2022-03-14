package models

import "time"

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
