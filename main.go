package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/user/mute/types"
)

const ConfigFilePath string = "config.json"

type Config struct {
	Accounts []Account
}

func ReadConfigFromJsonFile(path string) (Config, error) {
	var config Config

	b, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

type DB struct {
	// TODO: ...
}

func NewDB() (*DB, error) {
	// TODO: sql lite file setup
	return &DB{}, nil
}

func (db DB) InsertPost(sp SavedPost) error {
	// TODO: ...
	return nil
}

func (db DB) GetPostsByExternalID(externalID string) ([]SavedPost, error) {
	// TODO: ...
	return []SavedPost{}, nil
}

type Account struct {
	Name       string                         `json:"name"`
	ExternalID string                         `json:"externalId"`
	Creds      APICreds                       `json:"creds"`
	Sources    map[SourceName][]RssFeedSource `json:"sources"`
}

func (a *Account) GetSources() []Source {
	var sources []Source
	for _, srcs := range a.Sources {
		for _, s := range srcs {
			sources = append(sources, Source(s))
		}
	}
	return sources
}

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

type SourceName string

const (
	SourceNameRssFeed SourceName = "RssFeed"
)

type RssFeedSource struct {
	Url string
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

type Post struct {
	Url   string
	Title string
}

type SavedPost struct {
	ID string
	Post
}

func (p Post) ToSaved() SavedPost {
	return SavedPost{
		ID:   uuid.NewString(),
		Post: p,
	}
}

type APICreds struct {
	APIKey           string `json:"apiKey"`
	APIKeySecret     string `json:"apiKeySecret"`
	OAuthToken       string `json:"oAuthToken"`
	OAuthTokenSecret string `json:"oAuthTokenSecret"`
}

type TwitterClient struct {
	*gotwi.Client
	creds APICreds
}

func NewTwitterClient(apiKey, apiKeySecret, OAuthToken, OAuthTokenSecret string) (*TwitterClient, error) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		APIKey:               apiKey,
		APIKeySecret:         apiKeySecret,
		OAuthToken:           OAuthToken,
		OAuthTokenSecret:     OAuthTokenSecret,
	}
	c, err := gotwi.NewClient(in)
	if err != nil {
		return nil, err
	}
	creds := APICreds{
		APIKey:           apiKey,
		APIKeySecret:     apiKeySecret,
		OAuthToken:       OAuthToken,
		OAuthTokenSecret: OAuthTokenSecret,
	}
	twc := &TwitterClient{
		c,
		creds,
	}
	return twc, nil
}

type TwitterClientManager struct {
	Clients map[APICreds]*TwitterClient
}

func NewTwitterClientManager() TwitterClientManager {
	return TwitterClientManager{
		Clients: make(map[APICreds]*TwitterClient),
	}
}

func (tcm *TwitterClientManager) Get(creds APICreds) (*TwitterClient, error) {
	client, ok := tcm.Clients[creds]
	if ok {
		return client, nil
	}

	c, err := NewTwitterClient(creds.APIKey, creds.APIKeySecret, creds.OAuthToken, creds.OAuthTokenSecret)
	if err != nil {
		return nil, err
	}

	// Add the Twitter client to the map if login was successful
	tcm.Clients[creds] = c
	return c, nil
}

func (tcm *TwitterClientManager) PublishTweet(text string) (*types.CreateOutput, error) {
	// TODO: ...
	return nil, nil
}

func (tcm *TwitterClientManager) Remove(creds APICreds) bool {
	_, ok := tcm.Clients[creds]
	delete(tcm.Clients, creds)
	return ok
}

func PrintErr(err error) {
	timestamp := time.Now().Format(time.UnixDate)
	fmt.Printf("[%s] error: %s\n", timestamp, err.Error())
}

func PrintErrWithTimeout(err error, d time.Duration) {
	PrintErr(err)
	time.Sleep(d)
}

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	tcm := NewTwitterClientManager()

	for {
		config, err := ReadConfigFromJsonFile(ConfigFilePath)
		if err != nil {
			PrintErrWithTimeout(err, 10000)
		}

		fmt.Println(config)

		for _, acct := range config.Accounts {
			for _, source := range acct.GetSources() {
				posts, err := source.FetchPosts()
				if err != nil {
					PrintErrWithTimeout(err, 10000)
					continue
				}

				savedPosts, err := db.GetPostsByExternalID(acct.ExternalID)
				if err != nil {
					PrintErrWithTimeout(err, 10000)
					continue
				}

				// Compare posts with post history for this account
				var freshPosts []Post
				for _, post := range posts {
					for _, savedPost := range savedPosts {
						if post.Url != savedPost.Url {
							freshPosts = append(freshPosts, post)
						}
					}
				}

				for _, p := range freshPosts {
					text := source.FmtPost(p)
					if text == "" {
						PrintErrWithTimeout(fmt.Errorf("missing tweet text"), 10000)
						continue
					}

					_, err := tcm.PublishTweet(text)
					if err != nil {
						PrintErrWithTimeout(err, 10000)
						continue
					}

					if err := db.InsertPost(p.ToSaved()); err != nil {
						PrintErr(err)
					}
				}

			}
		}
	}
}
