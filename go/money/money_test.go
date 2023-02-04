package money_test

import (
	"money"
	"testing"
)

func TestMultiplication(t *testing.T) {
	t.Parallel()

	fiver := money.NewDollar(5)
	tenner := fiver.Times(2)

	if tenner.Amount() != 10 {
		t.Errorf("fiver.Times(2) = %d, want [%d]\n", tenner.Amount(), 10)
	}
}
