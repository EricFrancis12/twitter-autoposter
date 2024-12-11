package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/EricFrancis12/stripol"
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

	resp, err := http.Get(r.Url)
	if err != nil {
		return []Post{}, err
	}
	defer resp.Body.Close()

	// Read the entire response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Post{}, err
	}

	fmt.Println(string(body))

	return []Post{}, nil
}

func (r RssFeedSource) FmtPost(post Post) string {
	s := stripol.New("{{", "}}")

	s.RegisterVar("url", post.Url)
	s.RegisterVar("escapedUrl", url.PathEscape(post.Url))
	s.RegisterVar("title", post.Title)

	return s.Eval(r.TweetFmt)
}
