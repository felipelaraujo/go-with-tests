package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Em metódos que fazem alterações nos dados é preciso receber
// o ponteiro da struct que vai ser alterada.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Aqui não seria necessário, mas é bom colocar para manter
// a padronização do código.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
