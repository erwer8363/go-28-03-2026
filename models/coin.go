package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const API_URL = "https://api.coingecko.com/api/v3/simple/price"

type Coin struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type Price struct {
	Usd float64 `json:"usd"`
}

func (c *Coin) FetchPrice() (float64, error) {
	url := fmt.Sprintf("%s?vs_currencies=usd&ids=%s", API_URL, c.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("x-cg-demo-api-key", os.Getenv("API_KEY"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	price := map[string]Price{}
	json.Unmarshal(body, &price)
	fmt.Println("依次展示的吗。。。price:", price[c.ID], c.ID)
	return price[c.ID].Usd, nil
}

func (c *Coin) GetAmount() float64 {
	return c.Amount
}
