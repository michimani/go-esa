package models

import "time"

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
