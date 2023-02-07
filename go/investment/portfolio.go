package investment

// A Portfolio keeps track of zero or more Money pointers.
type Portfolio []*Money

// Add appends a *Money to a Portfolio.
func (p Portfolio) Add(m *Money) Portfolio {
	return append(p, m)
}

// Evaluate returns  the total value of a Portfolio in a specified currency.
func (p Portfolio) Evaluate(bank *Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make(ConversionErrors, 0, len(p))

	for _, m := range p {
		if converted, err := bank.Convert(m, currency); err == nil {
			total += converted.amount
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}

	if len(failedConversions) == 0 {
		return &Money{amount: total, currency: currency}, nil
	}

	return nil, failedConversions
}
