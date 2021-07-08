package pointerAndError

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	fmt.Println("address of balance in test is", &wallet.balance)
	want := Bitcoin(20)
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
