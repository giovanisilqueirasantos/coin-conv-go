package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	real, err := NewCurrency("real")
	assert.Nil(t, err)

	dollar, err := NewCurrency("dollar")
	assert.Nil(t, err)

	euro, err := NewCurrency("euro")
	assert.Nil(t, err)

	dollarAmount, err := NewAmount("10")
	assert.Nil(t, err)

	realAmount, err := NewAmount("10")
	assert.Nil(t, err)

	euroAmount, err := NewAmount("10")
	assert.Nil(t, err)

	btcAmount, err := NewAmount("10")
	assert.Nil(t, err)

	dollarToRealRate, err := NewAmount("4.50")
	assert.Nil(t, err)

	realToDollarRate, err := NewAmount("0.22")
	assert.Nil(t, err)

	realToEuroRate, err := NewAmount("0.19")
	assert.Nil(t, err)

	euroToRealRate, err := NewAmount("5.30")
	assert.Nil(t, err)

	btcToDollarRate, err := NewAmount("29932.60")
	assert.Nil(t, err)

	btcToRealRate, err := NewAmount("143112.32")
	assert.Nil(t, err)

	realToDollar, currencySymbol := ConvertCurrency(*dollar, *realAmount, *realToDollarRate)
	assert.Equal(t, amount(2.2), *realToDollar)
	assert.Equal(t, "$", currencySymbol)

	dollarToReal, currencySymbol := ConvertCurrency(*real, *dollarAmount, *dollarToRealRate)
	assert.Equal(t, amount(45), *dollarToReal)
	assert.Equal(t, "R$", currencySymbol)

	realToEuro, currencySymbol := ConvertCurrency(*euro, *realAmount, *realToEuroRate)
	assert.Equal(t, amount(1.9), *realToEuro)
	assert.Equal(t, "€", currencySymbol)

	euroToReal, currencySymbol := ConvertCurrency(*real, *euroAmount, *euroToRealRate)
	assert.Equal(t, amount(53), *euroToReal)
	assert.Equal(t, "R$", currencySymbol)

	btcToDollar, currencySymbol := ConvertCurrency(*dollar, *btcAmount, *btcToDollarRate)
	assert.Equal(t, amount(299326), *btcToDollar)
	assert.Equal(t, "$", currencySymbol)

	btcToReal, currencySymbol := ConvertCurrency(*real, *btcAmount, *btcToRealRate)
	assert.Equal(t, amount(1431123.2000000002), *btcToReal)
	assert.Equal(t, "R$", currencySymbol)
}
