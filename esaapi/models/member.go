package models

import "time"

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

type User struct {
	Myself     bool   `json:"myself"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type Me struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	ScreenName string     `json:"screen_name"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Icon       string     `json:"icon"`
	Email      string     `json:"email"`
	Teams      []Team     `json:"teams,omitempty"`
}
