package service

import (
	"context"

	"github.com/giovanisilqueirasantos/coin-conv-go/domain"
)

type Service interface {
	Exchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate domain.Amount) (*domain.Amount, error)
}

type service struct {
	Repo domain.Repo
}

func NewService(r domain.Repo) Service {
	return service{Repo: r}
}

func (s service) Exchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate domain.Amount) (*domain.Amount, error) {
	convAmount, err := s.Repo.GetExchange(ctx, currencyFrom, currencyTo, amountQuant, rate)
	if err != nil {
		return nil, err
	}

	if convAmount != nil {
		return convAmount, nil
	}

	newConvAmount := domain.ConvertCurrency(amountQuant, rate)

	err = s.Repo.StoreExchange(ctx, currencyFrom, currencyTo, amountQuant, rate, newConvAmount)
	if err != nil {
		return nil, err
	}

	return &newConvAmount, nil
}
