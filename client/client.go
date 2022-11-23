package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL_SERVER = "http://localhost:8080"

func main() {
	info, err := GetBid()
	if err != nil {
		panic(err)
	}

	fmt.Printf("BID value: %v \n", info)
}

func GetBid() (string, error) {
	res, err := http.Get(URL_SERVER)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
