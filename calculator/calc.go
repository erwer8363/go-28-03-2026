package calculator

import (
	"crypto-calc/api"
	"fmt"
)

func CalcTotal(fetchers []api.PriceFetcher) float64 {
	totalPrice := 0.0
	for _, fetcher := range fetchers {
		price, err := fetcher.FetchPrice()
		if err != nil {
			continue
		}
		fmt.Println("fetch price:", price)
		totalPrice += price * fetcher.GetAmount()
	}
	return totalPrice
}
