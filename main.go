package main

import (
	"errors"
	"fmt"
	"sync"
)

// Bitcoin type based on float64.
type Bitcoin float64

// Wallet is a type that allows the deposit
// and withdraw operations.
type Wallet struct {
	User    string
	Balance Bitcoin
	mutex   *sync.Mutex
}

// Withdraw method implements withdraw realisation.
func (w *Wallet) Withdraw(amount float64) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if w.Balance-Bitcoin(amount) < 0 {
		return errors.New("not enough money to withdraw")
	}
	w.Balance -= Bitcoin(amount)
	return nil
}

// Deposit method implements deposit realisation.
func (w *Wallet) Deposit(amount float64) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.Balance += Bitcoin(amount)
}

func (w *Wallet) string() string {
	return fmt.Sprintf("user %v balance is %g bitcoin \n", w.User, w.Balance)
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

// showMenu shows an interactive menu for the user.
func showMenu(w *Wallet) {
	for {
		var input int
		fmt.Printf("1. Amount \n")
		fmt.Printf("2. Deposit \n")
		fmt.Printf("3. Withdrow \n")
		fmt.Printf("4. Exit \n")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Incorrect input")
		}
		switch input {
		case 1:
			fmt.Println(w.string())
		case 2:
			var dep float64
			fmt.Printf("enter the deposit amount »")
			_, err := fmt.Scanln(&dep)
			if err != nil {
				fmt.Println("Incorrect input")
			}
			w.Deposit(dep)
			fmt.Printf("Depositing %g bitcoin \n", dep)
		case 3:
			var with float64
			fmt.Printf("enter the amount »")
			_, err := fmt.Scanln(&with)
			if err != nil {
				fmt.Println("Incorrect input")
			}
			err = w.Withdraw(with)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("This option doesn't exist")
		}
	}
}

func main() {
	var w = New("Nikita", 1.00)
	showMenu(w)
}
