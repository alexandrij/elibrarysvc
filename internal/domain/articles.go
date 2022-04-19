package domain

import "time"

type ArticleID uint64

type Article struct {
	ID      ArticleID
	Alias   string    `json:"alias,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Href    string    `json:"href,omitempty"`
	Author  string    `json:"author,omitempty"`
	Created time.Time `json:"created,omitempty"`
}
