package twitchclient

import (
	"fmt"
)

type ChatMessage struct {
	User      string
	Content   string
	Timestamp uint64
}

type ClientConfig struct {
	Username string
	Password string
	BaseUrl  string
	ApiKey   string
}

type Client struct {
	Username string
	Password string
	BaseUrl  string
	ApiKey   string
}

func Init(config ClientConfig) *Client {
	client := &Client{
		Username: config.Username,
		Password: config.Password,
		BaseUrl:  config.BaseUrl,
		ApiKey:   config.ApiKey,
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
