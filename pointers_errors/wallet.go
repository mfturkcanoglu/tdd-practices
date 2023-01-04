package main

import (
	"errors"
)

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var (
	ErrNotEnoughBalance = errors.New("not enough balance")
)

func (w *Wallet) WithDraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrNotEnoughBalance
	}

	w.balance -= amount
	return nil
}
