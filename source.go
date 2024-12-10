package main

import (
	"fmt"
)

type SourceName string

const (
	SourceNameRssFeed SourceName = "RssFeed"
)

type Source interface {
	GetSourceName() SourceName
	FetchPosts() ([]Post, error)
	FmtPost(Post) string
}

func NewSource(sourceName SourceName) Source {
	switch sourceName {
	case SourceNameRssFeed:
		return RssFeedSource{}
	}
	panic(fmt.Sprintf("unknown sourceName %s", sourceName))
}

type RssFeedSource struct {
	Url      string `json:"url"`
	TweetFmt string `json:"tweetFmt"`
}

func (r RssFeedSource) GetSourceName() SourceName {
	return SourceNameRssFeed
}

func (r RssFeedSource) FetchPosts() ([]Post, error) {
	// TODO: ...
	return []Post{}, nil
}

func (r RssFeedSource) FmtPost(post Post) string {
	// TODO: ...
	return ""
}
