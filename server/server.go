package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const urlAPI = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type USD_BRL struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/", GetUSDHandler)
	http.ListenAndServe(":8080", nil)
}

func GetUSDHandler(w http.ResponseWriter, r *http.Request) {
	usd, err := GetPrice()
	if err != nil {
		panic(err)
	}

	w.Write([]byte(usd.USDBRL.Bid))
	// json.NewEncoder(w).Encode(usd)
}

func GetPrice() (*USD_BRL, error) {
	res, err := http.Get(urlAPI)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Print(string(body))

	var usd USD_BRL
	err = json.Unmarshal(body, &usd)
	if err != nil {
		return nil, err
	}
	return &usd, nil
}
