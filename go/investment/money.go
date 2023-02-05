// Package investment implements types and functions for working with
// investments in various currencies.
package investment

// A Money stores information about money in any currency.
type Money struct {
	amount   float64
	currency string
}

// NewMoney returns a pointer to Money initialized to n dollars and c currency.
func NewMoney(n float64, c string) *Money {
	return &Money{
		amount:   n,
		currency: c,
	}
}

// Times returns a new pointer to Money initialized to d.amount times
// multiplier. The pointer's currency will be the same as that of m.
func (m *Money) Times(multiplier float64) *Money {
	return &Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// Divide returns a new pointer to Money initialized to m.amount divided by
// divisor. The pointer's currency will be the same as that of m.
func (m *Money) Divide(divisor float64) *Money {
	return &Money{
		amount:   m.amount / divisor,
		currency: m.currency,
	}
}

// Amount returns the current amount of currency stored in a Money.
func (m *Money) Amount() float64 {
	return m.amount
}
