package wallet

import (
	"errors"
	"fmt"
	"sync"
)

// NotEnoughMoneyToWithdraw custom error.
var NotEnoughMoneyToWithdraw = errors.New("not enough money to withdraw")

// IncorrectInput custom error.
var IncorrectInput = errors.New(" incorrect input")

// Bitcoin type based on float64.
type Bitcoin float64

// Wallet is a type that allows the deposit
// and withdraw operations.
type Wallet struct {
	User    string
	Balance Bitcoin
	mutex   *sync.Mutex
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.4f BTC", b)
}

func (w Wallet) String() string {
	return fmt.Sprintf("User %q balance is %s", w.User, w.Balance)
}

// Withdraw method implements withdraw realisation.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount <= 0 {
		return IncorrectInput
	}
	if w.Balance-amount < 0 {
		return NotEnoughMoneyToWithdraw
	}
	w.Balance -= amount
	return nil
}

// Deposit method implements deposit realisation.
func (w *Wallet) Deposit(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount <= 0 {
		return IncorrectInput
	}
	fmt.Printf("Depositing %s  \n", amount)
	w.Balance += amount
	return nil
}

// ShowBalance method return String method
// and shows menu.
func (w *Wallet) ShowBalance() string {
	return w.String()
}

// New function is a constructor for Wallet.
// New function returns wallet with initialized mutex
// and with some additional fields provided.
func New(name string, balance Bitcoin) *Wallet {
	return &Wallet{
		User:    name,
		Balance: balance,
		mutex:   new(sync.Mutex),
	}
}
