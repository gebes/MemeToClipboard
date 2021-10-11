package main

import (
	"encoding/json"
	"errors"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"net/http"
	"os"
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
		println("Could not fetch a meme:", err.Error())
		os.Exit(1)
	}

	err = clipboard.WriteAll(meme)
	if err != nil {
		println("Could not copy the meme to the clipboard:", err.Error())
		os.Exit(2)
	}

	println(meme)

}

func getMeme() (string, error){
	res, err := http.Get(Endpoint)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

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

	if len(response.Preview) == 0 {
		return "", errors.New("no meme found")
	}

	return response.Preview[len(response.Preview)-1], nil
}