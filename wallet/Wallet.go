package wallet

import "errors"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
