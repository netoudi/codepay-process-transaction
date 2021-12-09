package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "John Doe", 12, 2022, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "John Doe", 12, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardNumberExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "John Doe", 13, 2022, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "John Doe", 0, 2022, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "John Doe", 12, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardNumberExpirationYear(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "John Doe", 12, 2020, 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
