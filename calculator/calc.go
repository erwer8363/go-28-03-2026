package calculator

import (
	"crypto-calc/api"
	"fmt"
	"sync"
)

func CalcTotal(fetchers []api.PriceFetcher) float64 {
	totalPrice := 0.0
	var wg sync.WaitGroup
	for _, fetcher := range fetchers {
		wg.Add(1)
		go func() {
			price, err := fetcher.FetchPrice()
			if err != nil {
				return
			}
			fmt.Println("fetch price:", price)
			totalPrice += price * fetcher.GetAmount()
			wg.Done()
		}()
	}
	fmt.Println("这里显示了吗。。。totalPrice:", totalPrice)
	wg.Wait()
	return totalPrice
}
