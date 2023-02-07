package investment

import (
	"errors"
	"fmt"
)

// Bank stores exchange rates and methods for currency conversion.
type Bank struct {
	exchangeRates map[string]float64
}

// NewBank returns a pointer to an initialized but empty Bank.
func NewBank() *Bank {
	return &Bank{exchangeRates: make(map[string]float64)}
}

// AddExchangeRate creates or updates an exchange rate from one currency to another.
func (b *Bank) AddExchangeRate(from, to string, rate float64) {
	key := fmt.Sprintf("%s->%s", from, to)
	b.exchangeRates[key] = rate
}

// Convert attempts to convert a Money from one currency to another. If the
// conversion succeeds, the method returns a valid money in the new currency
// and a nil error. If the conversion fails, the method returns a nil Money and
// an error specifying the conversion that didn't work.
func (b *Bank) Convert(m *Money, to string) (*Money, error) {
	if m.currency == to {
		return NewMoney(m.amount, m.currency), nil
	}

	key := fmt.Sprintf("%s->%s", m.currency, to)
	rate, ok := b.exchangeRates[key]
	if !ok {
		return nil, errors.New(key)
	}

	return NewMoney(m.amount*rate, to), nil
}
