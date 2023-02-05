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
		total += convert(m, currency)
	}

	return &Money{amount: total, currency: currency}
}

func convert(m *Money, currency string) float64 {
	if m.currency == currency {
		return m.amount
	}

	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	key := m.currency + "->" + currency

	return exchangeRates[key] * m.amount
}
