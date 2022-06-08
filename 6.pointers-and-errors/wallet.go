package pointersanderrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// without pointer the reference to w is different because he is copy
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// this works: (*w).balance but:
	// the makers of Go deemed this notation cumbersome, 
	// so the language permits us to write  w.balance, without 
	// an explicit dereference. 
	//
	// These pointers to structs even have their own name: 
	// struct pointers and they are automatically dereferenced.
	//
	// ref https://golang.org/ref/spec#Method_values
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}