package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type Config struct {
	Accounts []Account
}

type Account struct {
	Name      string                         `json:"name"`
	TwitterID string                         `json:"twitterId"`
	Creds     APICreds                       `json:"creds"`
	Sources   map[SourceName][]RssFeedSource `json:"sources"`
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
