package domain

import "context"

type Repo interface {
	GetExchange(ctx context.Context, currencyFrom, currencyTo Currency, amountQuant, rate Amount) (*Amount, error)
	StoreExchange(ctx context.Context, currencyFrom, currencyTo Currency, amountQuant, rate, amountConv Amount) error
}
