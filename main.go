package main

import (
	"encoding/json"
	"errors"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	Endpoint = "https://meme-api.herokuapp.com/gimme/dankmemes"
)

type Response struct {
	Preview []string `json:"preview"`
}

func main() {

	meme, err := getMeme()
	if err != nil {
		log.Fatalln("Could not fetch a meme", err)
	}

	err = clipboard.WriteAll(meme)
	if err != nil {
		log.Fatalln("Could not copy the meme to the clipboard", err)
	}

	println("Copied meme to clipboard")
	println(meme)

}

func getMeme() (string, error){
	res, err := http.Get(Endpoint)
	defer res.Body.Close()

	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("status code " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Preview[len(response.Preview)-1], nil
}