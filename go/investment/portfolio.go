package investment

// A Portfolio keeps track of zero or more Money pointers.
type Portfolio []*Money

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
