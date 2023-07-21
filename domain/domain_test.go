package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	real, err := NewCurrency("real")
	assert.Nil(t, err)
	assert.NotNil(t, real)

	dollar, err := NewCurrency("dollar")
	assert.Nil(t, err)
	assert.NotNil(t, dollar)

	realAmount, err := NewAmount("10")
	assert.Nil(t, err)
	assert.NotNil(t, realAmount)

	dollarAmount, err := NewAmount("10")
	assert.Nil(t, err)
	assert.NotNil(t, dollarAmount)

	dollarRate, err := NewAmount("4.50")
	assert.Nil(t, err)
	assert.NotNil(t, dollarRate)

	realAmountConverted, currencySymbol := ConvertCurrency(*real, *dollarAmount, *dollarRate)
	assert.Equal(t, amount(45), *realAmountConverted)
	assert.Equal(t, "R$", currencySymbol)

	dolarAmountConverted, currencySymbol := ConvertCurrency(*dollar, *realAmount, *dollarRate)
	assert.Equal(t, amount(45), *realAmountConverted)
	assert.Equal(t, "$", currencySymbol)
}
