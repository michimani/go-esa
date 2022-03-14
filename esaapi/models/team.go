package models

type Team struct {
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
}
