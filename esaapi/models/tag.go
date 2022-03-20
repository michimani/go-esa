package models

type Tag struct {
	Name       string `json:"name"`
	PostsCount int    `json:"posts_count"`
}
