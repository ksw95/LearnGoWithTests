package wallet

import (
	"fmt"
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func assertError(t testing.TB, err error, expected error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != expected {
		t.Errorf("got %q want %q", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("did not want an error but got one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of balanace in test is %p \n", &wallet.balance)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}

		err := wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initialAmount := Bitcoin(12)

		wallet := Wallet{balance: initialAmount}

		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, initialAmount)
		assertError(t, err, ErrInsufficientFunds)

	})
}
