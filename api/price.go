package api

type PriceFetcher interface {
	FetchPrice() (float64, error)
	GetAmount() float64
}
