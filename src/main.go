package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	bot "teebee/src/bot"
	twitch "teebee/src/twitch"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getJoke(w http.ResponseWriter, r *http.Request) {
	clientConfig := twitch.ClientConfig{
		ClientId:     "",
		ClientSecret: "",
		BaseUrl:      "",
		ApiKey:       "",
	}
	client := twitch.Init(clientConfig)
	tBot := bot.Bot{Client: client}

	command := bot.Command{
		Command: bot.CmdJoke,
		Params:  make([]string, 0),
	}

	tBot.HandleCommand(command)
}

type Config struct {
	ClientId         string `json:"clientId"`
	ClientSecret     string `json:"clientSecret"`
	IRCAddress       string `json:"ircAddress"`
	WebsocketAddress string `json:"websocketAddress"`
}

func main() {
	configRaw, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	var config *Config
	json.Unmarshal(configRaw, &config)

	// getJoke(nil, nil)

	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", getHello)
	// http.HandleFunc("/joke", getJoke)

	// err := http.ListenAndServe(":3333", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }
}
