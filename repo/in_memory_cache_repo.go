package repo

import (
	"context"
	"fmt"
	"github.com/giovanisilqueirasantos/coin-conv-go/domain"
	"sync"
	"time"
)

type cacheItem struct {
	Value     domain.Amount
	ExpiresAt int64
}

type cache struct {
	Items map[string]cacheItem
	Stop  chan struct{}
	Wg    sync.WaitGroup
	Mu    sync.RWMutex
}

type inMemoryCacheRepo struct {
	Cache cache
	Repo  domain.Repo
}

func NewInMemoryCacheRepo(repo domain.Repo, cleanInterval time.Duration) *inMemoryCacheRepo {
	cache := cache{
		Items: make(map[string]cacheItem),
		Stop:  make(chan struct{}),
	}

	cache.Wg.Add(1)
	go func(cleanInterval time.Duration) {
		defer cache.Wg.Done()
		cache.cleanLoop(cleanInterval)
	}(cleanInterval)

	return &inMemoryCacheRepo{
		Cache: cache,
		Repo:  repo,
	}
}

func (r *cache) cleanLoop(cleanInterval time.Duration) {
	t := time.NewTicker(cleanInterval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			r.Mu.Lock()
			for key, item := range r.Items {
				if time.Now().Unix() > item.ExpiresAt {
					delete(r.Items, key)
				}
			}
			r.Mu.Unlock()
		case <-r.Stop:
			return
		}
	}
}

func (r *inMemoryCacheRepo) GetExchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate domain.Amount) (*domain.Amount, error) {
	if item, found := r.Cache.Items[currencyFrom.Text+currencyTo.Text+fmt.Sprintf("%f", amountQuant.Value)+fmt.Sprintf("%f", rate.Value)]; found {
		return &item.Value, nil
	}
	amount, err := r.Repo.GetExchange(ctx, currencyFrom, currencyTo, amountQuant, rate)
	if err != nil {
		return amount, err
	}
	if amount == nil {
		return nil, nil
	}
	item := cacheItem{Value: *amount, ExpiresAt: time.Now().Add(time.Hour).Unix()}
	r.Cache.Items[currencyFrom.Text+currencyTo.Text+fmt.Sprintf("%f", amountQuant.Value)+fmt.Sprintf("%f", rate.Value)] = item
	return amount, nil
}

func (r *inMemoryCacheRepo) StoreExchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate, convAmount domain.Amount) error {
	if err := r.Repo.StoreExchange(ctx, currencyFrom, currencyTo, amountQuant, rate, convAmount); err != nil {
		return err
	}
	item := cacheItem{Value: convAmount, ExpiresAt: time.Now().Add(time.Hour).Unix()}
	r.Cache.Items[currencyFrom.Text+currencyTo.Text+fmt.Sprintf("%f", amountQuant.Value)+fmt.Sprintf("%f", rate.Value)] = item
	return nil
}
