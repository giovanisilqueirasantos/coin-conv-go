package service

import (
	"context"

	"github.com/giovanisilqueirasantos/coin-conv-go/domain"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (ms MockService) Exchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate domain.Amount) (*domain.Amount, error) {
	args := ms.Called(ctx, currencyFrom, currencyTo, amountQuant, rate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	am := domain.Amount{}
	_ = am.New(args.String(0))
	return &am, nil
}
