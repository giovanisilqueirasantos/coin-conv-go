package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	real, err := NewCurrency("real")
	assert.Nil(t, err)
	assert.NotNil(t, real)
}
