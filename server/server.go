package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Print(string(body))

	var usd USD_BRL
	err = json.Unmarshal(body, &usd)
	if err != nil {
		return nil, err
	}

	InsertDb(usd)

	return &usd, nil
}

type USD struct {
	ID  string `json:id`
	Bid string `json:"bid"`
}

func NewUSD(usd USD_BRL) *USD {
	if usd.USDBRL.Bid == "" {
		panic("Bid not defined")
	}

	return &USD{
		ID:  uuid.New().String(),
		Bid: usd.USDBRL.Bid,
	}
}

func InsertDb(usdBrl USD_BRL) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/server_go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usd := NewUSD(usdBrl)

	stmt, err := db.Prepare("insert into prices(id, bid) values(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(usd.ID, usd.Bid)
	if err != nil {
		panic(err)
	}
}
