package main

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/user/mute/types"
)

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

func (tcm *TwitterClientManager) Remove(creds APICreds) bool {
	_, ok := tcm.Clients[creds]
	delete(tcm.Clients, creds)
	return ok
}

func (tcm *TwitterClientManager) PublishTweet(text string) (*types.CreateOutput, error) {
	// TODO: ...
	return nil, nil
}
