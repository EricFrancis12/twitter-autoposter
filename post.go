package main

import "github.com/google/uuid"

type Post struct {
	Url   string
	Title string
}

func NewPost(url string, title string) (*Post, error) {
	// Remove query string to ensure proper url match in autoposter
	url, err := stripQueryString(url)
	if err != nil {
		return nil, err
	}

	return &Post{
		Url:   url,
		Title: title,
	}, nil
}

func (p Post) ToSaved(twitterID string) SavedPost {
	return SavedPost{
		ID:        uuid.NewString(),
		TwitterID: twitterID,
		Post:      p,
	}
}

// TODO: refactor to combine In() and InSaved()

func (p Post) In(posts []Post) bool {
	return Some(posts, func(post Post) bool {
		return post.Url == p.Url
	})
}

func (p Post) InSaved(savedPosts []SavedPost) bool {
	return Some(savedPosts, func(savedPost SavedPost) bool {
		return savedPost.Url == p.Url
	})
}
