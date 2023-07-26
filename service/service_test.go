package service

import (
	"context"
	"errors"
	"testing"

	"github.com/giovanisilqueirasantos/coinconv/domain"
	"github.com/giovanisilqueirasantos/coinconv/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExchangeGetExchangeError(t *testing.T) {
	repo := repo.MockRepo{}

	currencyFrom := domain.Currency{}
	err := currencyFrom.New("real")
	assert.NoError(t, err)

	currencyTo := domain.Currency{}
	err = currencyTo.New("dollar")
	assert.NoError(t, err)

	amountQuant := domain.Amount{}
	err = amountQuant.New("10")
	assert.NoError(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.NoError(t, err)

	repo.On("GetExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate).Return(nil, errors.New("error message"))

	service := NewService(repo)

	_, err = service.Exchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate)

	assert.Error(t, err)
}

func TestExchangeStoreExchangeError(t *testing.T) {
	repo := repo.MockRepo{}

	currencyFrom := domain.Currency{}
	err := currencyFrom.New("real")
	assert.NoError(t, err)

	currencyTo := domain.Currency{}
	err = currencyTo.New("dollar")
	assert.NoError(t, err)

	amountQuant := domain.Amount{}
	err = amountQuant.New("10")
	assert.NoError(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.NoError(t, err)

	amountConv := domain.Amount{}
	err = amountConv.New("45")
	assert.NoError(t, err)

	repo.On("GetExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate).Return(nil, nil)

	repo.On("StoreExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate, amountConv).Return(errors.New("error message"))

	service := NewService(repo)

	_, err = service.Exchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate)

	assert.Error(t, err)
}

func TestExchangeGetExchange(t *testing.T) {
	repo := repo.MockRepo{}

	currencyFrom := domain.Currency{}
	err := currencyFrom.New("real")
	assert.NoError(t, err)

	currencyTo := domain.Currency{}
	err = currencyTo.New("dollar")
	assert.NoError(t, err)

	amountQuant := domain.Amount{}
	err = amountQuant.New("10")
	assert.NoError(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.NoError(t, err)

	repo.On("GetExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate).Return("45")

	service := NewService(repo)

	convAmount, err := service.Exchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate)

	assert.NoError(t, err)
	assert.NotNil(t, convAmount)
}

func TestExchange(t *testing.T) {
	repo := repo.MockRepo{}

	currencyFrom := domain.Currency{}
	err := currencyFrom.New("real")
	assert.NoError(t, err)

	currencyTo := domain.Currency{}
	err = currencyTo.New("dollar")
	assert.NoError(t, err)

	amountQuant := domain.Amount{}
	err = amountQuant.New("10")
	assert.NoError(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.NoError(t, err)

	amountConv := domain.Amount{}
	err = amountConv.New("45")
	assert.NoError(t, err)

	repo.On("GetExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate).Return(nil, nil)

	repo.On("StoreExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate, amountConv).Return(nil)

	service := NewService(repo)

	convAmount, err := service.Exchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate)

	assert.NoError(t, err)
	assert.NotNil(t, convAmount)
}
