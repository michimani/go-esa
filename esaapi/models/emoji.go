package models

type Emoji struct {
	Code     string   `json:"code"`
	Aliases  []string `json:"aliases"`
	Category string   `json:"category"`
	URL      string   `json:"url"`
}
