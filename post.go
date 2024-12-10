package main

import "github.com/google/uuid"

type Post struct {
	Url   string
	Title string
}

func (p Post) ToSaved() SavedPost {
	return SavedPost{
		ID:   uuid.NewString(),
		Post: p,
	}
}