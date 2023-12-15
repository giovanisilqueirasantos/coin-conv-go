package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/giovanisilqueirasantos/coin-conv-go/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCacheCleanLoop(t *testing.T) {
	repo := MockRepo{}

	cache := NewInMemoryCacheRepo(repo, time.Second)

	cache.Cache.Items["key"] = cacheItem{
		Value:     domain.Amount{},
		ExpiresAt: time.Now().Unix(),
	}

	time.Sleep(time.Second)
	assert.Equal(t, len(cache.Cache.Items), 0)
}

func TestGetExchangeFoundInCache(t *testing.T) {
	repo := MockRepo{}

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

	cache := NewInMemoryCacheRepo(repo, time.Hour)

	key := currencyFrom.Text + currencyTo.Text + fmt.Sprintf("%f", amountQuant.Value) + fmt.Sprintf("%f", rate.Value)

	cache.Cache.Items[key] = cacheItem{
		Value:     domain.Amount{},
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}

	amount, err := cache.GetExchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate)

	repo.AssertNotCalled(t, "GetExchange")
	assert.NoError(t, err)
	assert.Equal(t, amount.Value, cache.Cache.Items[key].Value.Value)
}

func TestGetExchangeNotFoundInCache(t *testing.T) {
	repo := MockRepo{}

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

	cache := NewInMemoryCacheRepo(repo, time.Hour)

	key := currencyFrom.Text + currencyTo.Text + fmt.Sprintf("%f", amountQuant.Value) + fmt.Sprintf("%f", rate.Value)

	amount, err := cache.GetExchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate)

	assert.NoError(t, err)
	assert.Equal(t, amount.Value, cache.Cache.Items[key].Value.Value)
}

func TestStoreExchangeInCache(t *testing.T) {
	repo := MockRepo{}

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

	convAmount := domain.Amount{}
	err = convAmount.New("45")
	assert.NoError(t, err)

	repo.On("StoreExchange", mock.Anything, currencyFrom, currencyTo, amountQuant, rate, convAmount).Return(nil, nil)

	cache := NewInMemoryCacheRepo(repo, time.Hour)

	key := currencyFrom.Text + currencyTo.Text + fmt.Sprintf("%f", amountQuant.Value) + fmt.Sprintf("%f", rate.Value)

	err = cache.StoreExchange(context.Background(), currencyFrom, currencyTo, amountQuant, rate, convAmount)

	assert.NoError(t, err)
	assert.Equal(t, convAmount.Value, cache.Cache.Items[key].Value.Value)
}
