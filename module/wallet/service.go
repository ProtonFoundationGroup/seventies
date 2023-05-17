package wallet

import "errors"

type Service struct {
	balance float64
}

func (w *Service) Deposit(amount float64) {
	w.balance += amount
}

func (w *Service) Withdraw(amount float64) error {
	if amount > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}

func (w *Service) Balance() float64 {
	return w.balance
}
