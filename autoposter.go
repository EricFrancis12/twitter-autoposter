package main

import (
	"fmt"
	"time"
)

type AutoPoster struct {
	configFilePath string
	timeout        time.Duration
	db             *DB
	tcm            TwitterClientManager
}

func NewAutoPoster(configFilePath string, timeout time.Duration) (*AutoPoster, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	return &AutoPoster{
		configFilePath: configFilePath,
		timeout:        timeout,
		db:             db,
		tcm:            NewTwitterClientManager(),
	}, nil
}

func (a *AutoPoster) Run() {
	for {
		config, err := ReadConfigFromJsonFile(ConfigFilePath)
		if err != nil {
			PrintErrWithTimeout(err, a.timeout)
		}

		fmt.Println(config)

		for _, acct := range config.Accounts {
			for _, source := range acct.GetSources() {
				posts, err := source.FetchPosts()
				if err != nil {
					PrintErrWithTimeout(err, a.timeout)
					continue
				}

				savedPosts, err := a.db.GetPostsByExternalID(acct.ExternalID)
				if err != nil {
					PrintErrWithTimeout(err, a.timeout)
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
						PrintErrWithTimeout(fmt.Errorf("missing tweet text"), a.timeout)
						continue
					}

					_, err := a.tcm.PublishTweet(text)
					if err != nil {
						PrintErrWithTimeout(err, a.timeout)
						continue
					}

					if err := a.db.InsertPost(p.ToSaved()); err != nil {
						PrintErr(err)
					}
				}

			}
		}
	}
}
