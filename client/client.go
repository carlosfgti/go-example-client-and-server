package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL_SERVER = "http://localhost:8080"

func main() {
	info, err := GetBid()
	if err != nil {
		panic(err)
	}

	fmt.Println(info)
}

func GetBid() (string, error) {
	res, err := http.Get(URL_SERVER)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
