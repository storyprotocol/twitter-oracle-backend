package twitter

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
	"net/http"
	"twitter-oracle-backend/logic/config"
)

var (
	client *twitter.Client
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func init() {
	twitterConfig := config.Get().Twitter
	client = newClientByOAuth1a(twitterConfig.ApiKey, twitterConfig.ApiSecret, twitterConfig.AccessToken, twitterConfig.AccessTokenSecret)
}

func GetClient() *twitter.Client {
	return client
}

func newClientByOAuth1a(apiKey, apiSecret, accessToken, accessTokenSecret string) *twitter.Client {
	oauth1Config := oauth1.NewConfig(apiKey, apiSecret)
	httpClient := oauth1Config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       accessToken,
		TokenSecret: accessTokenSecret,
	})

	return &twitter.Client{
		Authorizer: &authorize{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}
}

func newClientByToken(token string) *twitter.Client {
	return &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
}
