package money_test

import (
	"money"
	"testing"
)

func amountCheck(t *testing.T, s string, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("%s = [%+v], want [%+v]\n", s, got, want)
	}
}

func TestMultiplicationOfDollars(t *testing.T) {
	t.Parallel()

	fiver := money.NewMoney(5, "USD")
	tenner := fiver.Times(2)

	amountCheck(t, "fiver.Times(2)", tenner.Amount(), 10)
}

func TestMultiplicationOfEuros(t *testing.T) {
	t.Parallel()

	tenEuros := money.NewMoney(10, "EUR")
	twentyEuros := tenEuros.Times(2)

	amountCheck(t, "tenEuros.Times(2)", twentyEuros.Amount(), 20)
}

func TestDivisionOfWon(t *testing.T) {
	t.Parallel()

	m := money.NewMoney(4002, "KRW")
	got := m.Divide(4)
	want := money.NewMoney(1000.5, "KRW")

	amountCheck(t, "m.Divide(4)", got.Amount(), want.Amount())
}

func TestAddition(t *testing.T) {
	t.Parallel()
	var portfolio money.Portfolio
	var portfolioInDollars *money.Money

	fiveDollars := money.NewMoney(5, "USD")
	tenDollars := money.NewMoney(10, "USD")
	fifteenDollars := money.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	amountCheck(
		t,
		"portfolio.Add(fiveDollars) + portfolio.Add(tenDollars)",
		portfolioInDollars.Amount(),
		fifteenDollars.Amount(),
	)
}
