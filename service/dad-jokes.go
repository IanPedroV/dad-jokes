package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetJokeId() string {
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	stringBody := string(body)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(stringBody), &data)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", data["id"])
}

func GetDadJokeImage(jokeId string) string {
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
