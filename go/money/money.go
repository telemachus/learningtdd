// Package money implements arithmetic and conversion for several currencies.
package money

// A Dollar stores information about money in US dollars.
type Dollar struct {
	amount int
}

// NewDollar returns a pointer to Dollar initialized to n dollars.
func NewDollar(n int) *Dollar {
	return &Dollar{amount: n}
}

// Times returns a new pointer to Dollar initialized to d.amount times multiplier.
func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{amount: d.amount * multiplier}
}

// Amount returns the current number of dollars stored in a Dollar.
func (d *Dollar) Amount() int {
	return d.amount
}
