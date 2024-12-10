package main

import "github.com/google/uuid"

type Post struct {
	Url   string
	Title string
}

func (p Post) ToSaved(twitterID string) SavedPost {
	return SavedPost{
		ID:        uuid.NewString(),
		TwitterID: twitterID,
		Post:      p,
	}
}
