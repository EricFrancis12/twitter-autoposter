package main

import "github.com/google/uuid"

type Post struct {
	Url   string
	Title string
}

func NewPost(url string, title string) *Post {
	return &Post{
		Url:   url,
		Title: title,
	}
}

func (p Post) ToSaved(twitterID string) SavedPost {
	return SavedPost{
		ID:        uuid.NewString(),
		TwitterID: twitterID,
		Post:      p,
	}
}
