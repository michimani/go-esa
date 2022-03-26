package models

import "time"

type EmailInvitations struct {
	Email     string     `json:"email"`
	Code      string     `json:"code"`
	ExpiresAt *time.Time `json:"expires_at"`
	URL       string     `json:"url"`
}
