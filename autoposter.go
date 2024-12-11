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
	db, err := NewDB(DriverName, DataSourceName)
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
		PrintWithTimestamp("Reading config file")

		config, err := ReadConfigFromJsonFile(ConfigFilePath)
		if err != nil {
			PrintErrWithTimeout(err, a.timeout)
		}

		PrintWithTimestampf("Starting range over %d Accounts", len(config.Accounts))
		for _, acct := range config.Accounts {
			sources := acct.GetSources()

			PrintWithTimestampf("Starting range over %d Sources", len(sources))
			for _, source := range sources {
				posts, err := source.FetchPosts()
				if err != nil {
					PrintErrWithTimeout(err, a.timeout)
					continue
				}

				savedPosts, err := a.db.GetSavedPostsByTwitterID(acct.TwitterID)
				if err != nil {
					PrintErrWithTimeout(err, a.timeout)
					continue
				}

				// Compare posts with post history for this account
				var freshPosts []Post
				if len(savedPosts) <= 0 {
					freshPosts = posts
				} else {
					for _, post := range posts {
						if !post.InSaved(savedPosts) {
							freshPosts = append(freshPosts, post)
						}
					}
				}

				PrintWithTimestampf("Starting range over %d Fresh Posts", len(freshPosts))
				for _, p := range freshPosts {
					text := source.FmtPost(p)
					if text == "" {
						PrintErrWithTimeout(fmt.Errorf("missing tweet text"), a.timeout)
						continue
					}

					PrintWithTimestampf("Publishing tweet to %s: %s", acct.Name, text)
					_, err := a.tcm.PublishTweet(acct.Creds, text)
					if err != nil {
						PrintErrWithTimeout(err, a.timeout)
						continue
					}

					if err := a.db.InsertSavedPost(p.ToSaved(acct.TwitterID)); err != nil {
						PrintErr(err)
					}

					Sleep(a.timeout)
				}

			}
		}

		Sleep(a.timeout)
	}
}
