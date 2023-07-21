package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

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

type amount float64

func NewAmount(text string) (*amount, error) {
	if text == "" {
		return nil, errors.New("text can not be empty")
	} else if strings.Contains(text, ".") {
		if !regexp.MustCompile(`^(?:[0-9])[0-9]*\.(?:[0-9])[0-9]*$`).MatchString(text) {
			return nil, errors.New("text in worng format n.n")
		}
	} else {
		if !regexp.MustCompile(`^[0-9]*$`).MatchString(text) {
			return nil, errors.New("text can only contain numbers")
		}
	}

	textSplited := strings.Split(text, ".")

	if len(textSplited) > 1 {
		if len(textSplited[1]) > 2 {
			textSplited[1] = textSplited[1][0:2]
		}
	}

	textFloat, err := strconv.ParseFloat(strings.Join(textSplited, "."), 64)
	if err != nil {
		return nil, err
	}

	am := amount(textFloat)
	return &am, nil
}

func ConvertCurrency(to currency, amountQuant amount, rate amount) (*amount, string) {
	am := amount(rate * amountQuant)
	return &am, to.Symbol
}
