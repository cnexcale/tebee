package twitchclient

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

type ChatMessage struct {
	User      string
	Content   string
	Timestamp uint64
}

type ClientConfig struct {
	ClientId     string
	ClientSecret string
	BaseUrl      string
	ApiKey       string
}

type Client struct {
	Username string
	Password string
	BaseUrl  string
	ApiToken string
}

func Init(config ClientConfig) *Client {
	// https://github.com/twitchdev/authentication-go-sample
	oauth2Config := &clientcredentials.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		Username: config.ClientId,
		Password: config.ClientSecret,
		BaseUrl:  config.BaseUrl,
		ApiToken: token.AccessToken,
	}

	fmt.Println("Initialized client")

	return client
}

func (c Client) ReceiveMessage() ChatMessage {
	m := ChatMessage{
		User:      "TestUser",
		Content:   "lorem ipsum",
		Timestamp: 1000000,
	}

	// poll for message
	// parse message
	// return ChatMessage

	return m
}

func (c Client) SendMessage(message string) {

}
