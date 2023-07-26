package repo

import (
	"context"

	"github.com/giovanisilqueirasantos/coinconv/domain"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (mr MockRepo) GetExchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate domain.Amount) (*domain.Amount, error) {
	args := mr.Called(ctx, currencyFrom, currencyTo, amountQuant, rate)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Error(1)
	}

	am := domain.Amount{}
	_ = am.New(args.String(0))
	return &am, nil
}

func (mr MockRepo) StoreExchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate, amountConv domain.Amount) error {
	args := mr.Called(ctx, currencyFrom, currencyTo, amountQuant, rate, amountConv)
	return args.Error(0)
}
