package investment_test

import (
	"investment"
	"testing"
)

func amountCheck(t *testing.T, s string, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("%s = [%+v]; want [%+v]\n", s, got, want)
	}
}

func TestMultiplication(t *testing.T) {
	t.Parallel()

	tenEuros := investment.NewMoney(10, "EUR")
	twentyEuros := tenEuros.Times(2)

	amountCheck(t, "tenEuros.Times(2)", twentyEuros.Amount(), 20)
}

func TestDivision(t *testing.T) {
	t.Parallel()

	m := investment.NewMoney(4002, "KRW")
	got := m.Divide(4)
	want := investment.NewMoney(1000.5, "KRW")

	amountCheck(t, "m.Divide(4)", got.Amount(), want.Amount())
}

func TestAddition(t *testing.T) {
	t.Parallel()

	var portfolio investment.Portfolio
	var portfolioInDollars *investment.Money

	fiveDollars := investment.NewMoney(5, "USD")
	tenDollars := investment.NewMoney(10, "USD")
	fifteenDollars := investment.NewMoney(15, "USD")

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

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	t.Parallel()

	var portfolio investment.Portfolio

	fiveDollars := investment.NewMoney(5, "USD")
	tenEuros := investment.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	got := portfolio.Evaluate("USD")
	want := investment.NewMoney(17, "USD")

	amountCheck(
		t,
		"portfolio.Add(fiveDollars) + portfolio.Add(tenEuros)",
		got.Amount(),
		want.Amount(),
	)
}
