package models

import "time"

type Watcher struct {
	CreatedAt *time.Time `json:"created_at"`
	User      User       `json:"user"`
}
