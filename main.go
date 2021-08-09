package main

import (
	"fmt"
	"github.com/Nikby53/bitcoin-wallet/wallet"
)

// showMenu shows an interactive menu for the user.
func showMenu(w *wallet.Wallet) {
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
			fmt.Println(w.ShowBalance())
		case 2:
			var dep wallet.Bitcoin
			fmt.Printf("enter the deposit amount »")
			_, err := fmt.Scanln(&dep)
			if err != nil {
				fmt.Println("Incorrect input")
			}
			w.Deposit(dep)

		case 3:
			var with wallet.Bitcoin
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
	var w = wallet.New("Nikita", 0.00)
	showMenu(w)
}
