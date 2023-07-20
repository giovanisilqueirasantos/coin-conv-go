package domain

import "errors"

type currency struct {
	Text   string
	Symbol string
}

var validCurrencies = []currency{
	{Text: "real", Symbol: "R$"},
	{Text: "dollar", Symbol: "$"},
	{Text: "euro", Symbol: "€"},
	{Text: "btc", Symbol: "₿"},
}

func NewCurrency(text string) (*currency, error) {
	var cur *currency

	for _, c := range validCurrencies {
		if c.Text == text {
			cur = &currency{Text: c.Text, Symbol: c.Symbol}
		}
	}

	if cur == nil {
		return nil, errors.New("invalid currency text")
	}

	return cur, nil
}
