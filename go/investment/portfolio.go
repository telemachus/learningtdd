package investment

import (
	"errors"
	"strings"
)

// A Portfolio keeps track of zero or more Money pointers.
type Portfolio []*Money

// Add appends a *Money to a Portfolio.
func (p Portfolio) Add(m *Money) Portfolio {
	return append(p, m)
}

// Evaluate returns  the total value of a Portfolio in a specified currency.
func (p Portfolio) Evaluate(currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)

	for _, m := range p {
		if convertedAmount, ok := convert(m, currency); ok {
			total += convertedAmount
		} else {
			failure := m.currency + "->" + currency
			failedConversions = append(failedConversions, failure)
		}
	}

	if len(failedConversions) == 0 {
		return &Money{amount: total, currency: currency}, nil
	}

	failures := strings.Join(failedConversions, ", ")
	failures = "[" + failures + "]"
	err := errors.New("Missing exchange rate(s): " + failures)

	return &Money{}, err
}

func convert(m *Money, currency string) (float64, bool) {
	if m.currency == currency {
		return m.amount, true
	}

	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	key := m.currency + "->" + currency
	rate, ok := exchangeRates[key]

	return rate * m.amount, ok
}
