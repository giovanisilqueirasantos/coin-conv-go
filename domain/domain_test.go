package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	real := Currency{}
	err := real.New("real")
	assert.Nil(t, err)

	dollar := Currency{}
	err = dollar.New("dollar")
	assert.Nil(t, err)

	euro := Currency{}
	err = euro.New("euro")
	assert.Nil(t, err)

	dollarAmount := Amount{}
	err = dollarAmount.New("10")
	assert.Nil(t, err)

	realAmount := Amount{}
	err = realAmount.New("10")
	assert.Nil(t, err)

	euroAmount := Amount{}
	err = euroAmount.New("10")
	assert.Nil(t, err)

	btcAmount := Amount{}
	err = btcAmount.New("10")
	assert.Nil(t, err)

	dollarToRealRate := Amount{}
	err = dollarToRealRate.New("4.50")
	assert.Nil(t, err)

	realToDollarRate := Amount{}
	err = realToDollarRate.New("0.22")
	assert.Nil(t, err)

	realToEuroRate := Amount{}
	err = realToEuroRate.New("0.19")
	assert.Nil(t, err)

	euroToRealRate := Amount{}
	err = euroToRealRate.New("5.30")
	assert.Nil(t, err)

	btcToDollarRate := Amount{}
	err = btcToDollarRate.New("29932.60")
	assert.Nil(t, err)

	btcToRealRate := Amount{}
	err = btcToRealRate.New("143112.32")
	assert.Nil(t, err)

	realToDollar, currencySymbol := ConvertCurrency(dollar, realAmount, realToDollarRate)
	assert.Equal(t, 2.2, realToDollar.Value)
	assert.Equal(t, "$", currencySymbol)

	dollarToReal, currencySymbol := ConvertCurrency(real, dollarAmount, dollarToRealRate)
	assert.Equal(t, float64(45), dollarToReal.Value)
	assert.Equal(t, "R$", currencySymbol)

	realToEuro, currencySymbol := ConvertCurrency(euro, realAmount, realToEuroRate)
	assert.Equal(t, 1.9, realToEuro.Value)
	assert.Equal(t, "€", currencySymbol)

	euroToReal, currencySymbol := ConvertCurrency(real, euroAmount, euroToRealRate)
	assert.Equal(t, float64(53), euroToReal.Value)
	assert.Equal(t, "R$", currencySymbol)

	btcToDollar, currencySymbol := ConvertCurrency(dollar, btcAmount, btcToDollarRate)
	assert.Equal(t, float64(299326), btcToDollar.Value)
	assert.Equal(t, "$", currencySymbol)

	btcToReal, currencySymbol := ConvertCurrency(real, btcAmount, btcToRealRate)
	assert.Equal(t, 1431123.2000000002, btcToReal.Value)
	assert.Equal(t, "R$", currencySymbol)
}
