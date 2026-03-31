package main

import (
	"crypto-calc/api"
	"crypto-calc/calculator"
	"crypto-calc/models"
	"fmt"
	"path"
	"time"
)

import "github.com/joho/godotenv"

func main() {
	godotenv.Load(path.Join(".", ".env"))
	coinList := [3]models.Coin{
		{
			ID:     "bitcoin",
			Name:   "Bitcoin",
			Symbol: "btc",
			Amount: 1,
		},
		{
			ID:     "ethereum",
			Name:   "Ethereum",
			Symbol: "eth",
			Amount: 10,
		},
		{
			ID:     "binancecoin",
			Name:   "BNB",
			Symbol: "bnb",
			Amount: 20,
		},
	}

	coins := make([]api.PriceFetcher, 0)

	for _, coin := range coinList {
		coins = append(coins, &coin)
	}
	newTotal := calculator.CalcTotal(coins)
	fmt.Println("Hello ever", newTotal)
	time.Sleep(5 * time.Second)
}
