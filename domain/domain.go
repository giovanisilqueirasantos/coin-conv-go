package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Currency struct {
	Text   string
	Symbol string
}

var validCurrencies = []Currency{
	{Text: "real", Symbol: "R$"},
	{Text: "dollar", Symbol: "$"},
	{Text: "euro", Symbol: "€"},
	{Text: "btc", Symbol: "₿"},
}

func (cur *Currency) New(text string) error {
	for _, c := range validCurrencies {
		if c.Text == text {
			cur.Text = c.Text
			cur.Symbol = c.Symbol
			return nil
		}
	}

	return errors.New("invalid currency text")
}

func (cur Currency) Valid() bool {
	for _, c := range validCurrencies {
		if c.Text == cur.Text && c.Symbol == cur.Symbol {
			return true
		}
	}

	return false
}

type Amount struct {
	Value float64
}

func (a *Amount) New(text string) error {
	if text == "" {
		return errors.New("text can not be empty")
	} else if strings.Contains(text, ".") {
		if !regexp.MustCompile(`^(?:[0-9])[0-9]*\.(?:[0-9])[0-9]*$`).MatchString(text) {
			return errors.New("text in worng format n.n")
		}
	} else {
		if !regexp.MustCompile(`^[0-9]*$`).MatchString(text) {
			return errors.New("text can only contain numbers")
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
		return err
	}

	a.Value = textFloat
	return nil
}

func ConvertCurrency(amountQuant, rate Amount) Amount {
	return Amount{Value: rate.Value * amountQuant.Value}
}
