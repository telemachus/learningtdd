// Package money implements arithmetic and conversion for several currencies.
package money

// A Money stores information about money in any currency.
type Money struct {
	amount   float64
	currency string
}

// A Portfolio keeps track of zero or more Money pointers.
type Portfolio []*Money

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

// Add appends a *Money to a Portfolio.
func (p Portfolio) Add(m *Money) Portfolio {
	return append(p, m)
}

// Evaluate returns  the total value of a Portfolio in a specified currency.
func (p Portfolio) Evaluate(currency string) *Money {
	total := 0.0
	for _, m := range p {
		total += m.amount
	}

	return &Money{amount: total, currency: currency}
}
