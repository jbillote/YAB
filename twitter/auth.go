package twitter

import (
	"context"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/dghubble/go-twitter/twitter"

	"github.com/jbillote/YAB/config"
)

func TwitterAuth() *twitter.Client {
	log.Info("Creating Twitter client")

	c := config.GetConfig()
	oauthConfig := &clientcredentials.Config{
		ClientID: c.TwitterToken,
		ClientSecret: c.TwitterSecret,
		TokenURL: "https://api.twitter.com/oauth2/token",
	}

	httpClient := oauthConfig.Client(context.TODO())

	client := twitter.NewClient(httpClient)

	return client
}
