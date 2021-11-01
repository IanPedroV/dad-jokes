package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	Id     string  `json:"id"`
	Joke   string  `json:"joke"`
	Status float64 `json:"status"`
}

func Get() Joke {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"Accept": []string{"application/json"},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	joke := Joke{}
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		panic(err)
	}
	return joke
}

func GetImage(jokeId string) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/j/"+jokeId+".png", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Accept": []string{"text/plain"},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
