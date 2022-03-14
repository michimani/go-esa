package models

import "time"

type Stargazer struct {
	CreatedAt *time.Time `json:"created_at"`
	Body      string
	User      User
}
