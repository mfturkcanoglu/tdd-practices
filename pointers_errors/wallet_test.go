package main

import "testing"

func TestWaller(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := &Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := &Wallet{balance: Bitcoin(20)}
		err := wallet.WithDraw(Bitcoin(15))

		want := Bitcoin(5)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := &Wallet{balance: startingBalance}
		err := wallet.WithDraw(Bitcoin(100))

		assertError(t, err, ErrNotEnoughBalance)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t *testing.T, wallet *Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("Expected %s but actually got %s", want, got)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted error but did not get one")
	}
	if err != want {
		t.Errorf("Expected %q message but got %q", want.Error(), err.Error())
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but did not want one")
	}
}
