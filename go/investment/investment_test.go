package investment_test

import (
	"investment"
	"testing"
)

func getBank() *investment.Bank {
	b := investment.NewBank()
	b.AddExchangeRate("EUR", "USD", 1.2)
	b.AddExchangeRate("USD", "KRW", 1100)
	b.AddExchangeRate("EUR", "KRW", 1100)

	return b
}

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

	fiveDollars := investment.NewMoney(5, "USD")
	tenDollars := investment.NewMoney(10, "USD")
	fifteenDollars := investment.NewMoney(15, "USD")

	var portfolio investment.Portfolio
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	var portfolioInDollars *investment.Money
	portfolioInDollars, err := portfolio.Evaluate(getBank(), "USD")
	if err != nil {
		t.Fatalf("portfolio.Add(fiveDollars) + portfolio.Add(tenDollars) has error [%+v]; want nil", err)
	}

	amountCheck(
		t,
		"portfolio.Add(fiveDollars) + portfolio.Add(tenDollars)",
		portfolioInDollars.Amount(),
		fifteenDollars.Amount(),
	)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	t.Parallel()

	fiveDollars := investment.NewMoney(5, "USD")
	tenEuros := investment.NewMoney(10, "EUR")

	var portfolio investment.Portfolio
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	got, err := portfolio.Evaluate(getBank(), "USD")
	want := investment.NewMoney(17, "USD")

	if err != nil {
		t.Fatalf("portfolio.Add(fiveDollars) + portfolio.Add(tenEuros) has error: [%+v]; want nil", err)
	}

	amountCheck(
		t,
		"portfolio.Add(fiveDollars) + portfolio.Add(tenEuros)",
		got.Amount(),
		want.Amount(),
	)
}

func TestAdditionOfDollarsAndWon(t *testing.T) {
	t.Parallel()

	oneDollar := investment.NewMoney(1, "USD")
	elevenHundredWon := investment.NewMoney(1100, "KRW")

	var portfolio investment.Portfolio
	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)

	got, err := portfolio.Evaluate(getBank(), "KRW")
	want := investment.NewMoney(2200, "KRW")

	if err != nil {
		t.Fatalf("portfolio.Add(oneDollar) + portfolio.Add(elevenHundredWon) has error [%+v]; want nil", err)
	}

	amountCheck(
		t,
		"portfolio.Add(fiveDollars) + portfolio.Add(tenEuros)",
		got.Amount(),
		want.Amount(),
	)
}

func TestSuccessfulConvert(t *testing.T) {
	t.Parallel()

	bank := investment.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	tenEuros := investment.NewMoney(10, "EUR")
	got, err := bank.Convert(tenEuros, "USD")
	want := investment.NewMoney(12, "USD")

	if err != nil {
		t.Fatalf(
			"bank.Convert(tenEuros, \"USD\") error: [%+v]; want nil",
			err,
		)
	}

	amountCheck(
		t,
		"bank.Convert(tenEuros, \"USD\")",
		got.Amount(),
		want.Amount(),
	)
}

func TestUnsuccessfulConvert(t *testing.T) {
	t.Parallel()

	bank := investment.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	tenEuros := investment.NewMoney(10, "EUR")
	_, err := bank.Convert(tenEuros, "KAL")

	if err == nil {
		t.Error("bank.Convert(tenEuros, \"KAL\") is nil; want error")
	}
}

func TestSuccessfulMultiCurrencyEvaluation(t *testing.T) {
	t.Parallel()

	oneDollar := investment.NewMoney(1, "USD")
	oneEuro := investment.NewMoney(1, "EUR")
	elevenHundredWon := investment.NewMoney(1100, "KRW")

	var portfolio investment.Portfolio
	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(elevenHundredWon)

	want := 3300.0
	got, err := portfolio.Evaluate(getBank(), "KRW")
	if err != nil {
		t.Fatal("portfolio.Evaluate(getBank(), \"KRW\") error should be nil")
	}

	amountCheck(
		t,
		"portfolio.Evaluate(getBank(), \"KRW\")",
		got.Amount(),
		want,
	)
}

func TestUnsuccessfulMultiCurrencyEvaluation(t *testing.T) {
	t.Parallel()

	var portfolio investment.Portfolio
	oneDollar := investment.NewMoney(1, "USD")
	portfolio = portfolio.Add(oneDollar)

	got, err := portfolio.Evaluate(getBank(), "FOO")

	if got != nil {
		t.Error("portfolio.Evaluate(getBank(), \"FOO\") *Money should be nil")
	}

	if err == nil {
		t.Error("portfolio.Evaluate(getBank(), \"FOO\") error should not be nil")
	}
}
