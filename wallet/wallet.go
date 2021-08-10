package wallet

import (
	"errors"
	"fmt"
	"sync"
)

var (
	// ErrInsufficientFunds custom error.
	ErrInsufficientFunds = errors.New("insufficient funds")
	// ErrIncorrectInput custom error.
	ErrIncorrectInput = errors.New("incorrect input")
)

// Bitcoin type based on float64.
type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.4f BTC", b)
}

// Wallet is a type that allows the deposit, balance
// and withdraw operations.
type Wallet struct {
	user    string
	balance Bitcoin
	mutex   sync.RWMutex
}

func (w *Wallet) String() string {
	return fmt.Sprintf("User %q balance is %s", w.user, w.balance)
}

// Withdraw method implements withdraw realisation.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount <= 0 {
		return ErrIncorrectInput
	}
	if w.balance-amount < 0 {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Deposit method implements deposit realisation.
func (w *Wallet) Deposit(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount <= 0 {
		return ErrIncorrectInput
	}
	w.balance += amount
	return nil
}

// Balance method returns a String method .
func (w *Wallet) Balance() Bitcoin {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	return w.balance
}

// New function is a constructor for Wallet.
// New function returns wallet with initialized mutex
// and with some additional fields provided.
func New(name string, balance Bitcoin) *Wallet {
	return &Wallet{
		user:    name,
		balance: balance,
	}
}
