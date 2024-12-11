package main

import (
	"fmt"
	"net/url"

	"github.com/EricFrancis12/stripol"
	"github.com/mmcdole/gofeed"
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
	var posts []Post

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(r.Url)
	if err != nil {
		return posts, err
	}

	for _, item := range feed.Items {
		posts = append(posts, *NewPost(item.Link, item.Title))
	}

	return posts, nil
}

func (r RssFeedSource) FmtPost(post Post) string {
	s := stripol.New("{{", "}}")

	s.RegisterVar("url", post.Url)
	s.RegisterVar("escapedUrl", url.PathEscape(post.Url))
	s.RegisterVar("title", post.Title)

	return s.Eval(r.TweetFmt)
}
