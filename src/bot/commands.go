package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	http "net/http"

	twitch "teebee/src/twitch"
)

// https://jokeapi.dev/
const JokeApiUrl = "https://v2.jokeapi.dev/joke/Any?lang=de&blacklistFlags=nsfw,religious,political,racist,sexist,explicit"
const JokeApiUrlContainsParam = "&contains="

//	{
//	    "error": false,
//	    "category": "Programming",
//	    "type": "single",
//	    "joke": "Die Selbsthilfegruppe \"HTML-Sonderzeichen-Probleme\" trifft sich heute im gro&szlig;en Saal.",
//	    "flags": {
//	        "nsfw": false,
//	        "racist": false,
//	        "sexist": false,
//	        "religious": false,
//	        "political": false,
//	        "explicit": false
//	    },
//	    "id": 14,
//	    "safe": true,
//	    "lang": "de"
//	}
type JokeApiResponse struct {
	Error    string `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Safe     bool   `json:"safe"`
}

func handleJokeCommand(client twitch.Client, command Command) {
	response, err := http.Get(JokeApiUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var joke *JokeApiResponse

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, &joke)

	if joke.Type == "single" {
		fmt.Println(joke.Joke)
	} else {
		fmt.Println(joke.Setup + " - " + joke.Delivery)
	}
}
