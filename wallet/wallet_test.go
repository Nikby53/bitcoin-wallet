package wallet

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestWallet_Deposit(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          Bitcoin
		deposit       Bitcoin
		name          string
		expectedError error
	}{
		{
			wallet:  New("Nikita", 1.00),
			want:    5.32,
			deposit: 4.32,
			name:    "success",
		},
		{
			wallet:        New("Nikita", 1.00),
			want:          1.00,
			expectedError: ErrIncorrectInput,
			deposit:       -1.00,
			name:          "fail: incorrect input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wallet.Deposit(tt.deposit)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
			if tt.wallet.balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.balance)
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          Bitcoin
		withdraw      Bitcoin
		name          string
		expectedError error
	}{
		{
			wallet:   New("Nikita", 1.00),
			want:     0.50,
			withdraw: 0.50,
			name:     "success",
		},
		{
			wallet:        New("Nikita", 1.00),
			expectedError: ErrInsufficientFunds,
			withdraw:      2,
			name:          "fail:not enough money to withdraw",
			want:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wallet.Withdraw(tt.withdraw)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
			if tt.wallet.balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.balance)
			}
		})
	}
}

func TestWallet_WithdrawConcurrent(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          Bitcoin
		withdraw      Bitcoin
		name          string
		expectedError error
	}{
		{
			wallet:   New("Nikita", 5.00),
			want:     2.00,
			withdraw: 1.00,
			name:     "success",
		},
		{
			wallet:        New("Nikita", 1.00),
			expectedError: ErrInsufficientFunds,
			withdraw:      2.00,
			name:          "fail:not enough money to withdraw",
			want:          1.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			for i := 0; i < 3; i++ {
				wg.Add(1)
				go func() {
					err := tt.wallet.Withdraw(tt.withdraw)
					if !errors.Is(err, tt.expectedError) {
						t.Errorf("expected %v instead of %v", tt.expectedError, err)
					}
					defer wg.Done()
				}()
			}
			wg.Wait()
			if tt.wallet.balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.balance)
			}
		})
	}
}

func TestWallet_DepositConcurrent(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          Bitcoin
		deposit       Bitcoin
		name          string
		expectedError error
	}{
		{
			wallet:        New("Nikita", 1.00),
			want:          7.00,
			deposit:       2.00,
			expectedError: ErrIncorrectInput,
			name:          "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			for i := 0; i < 3; i++ {
				wg.Add(1)
				go func() {
					err := tt.wallet.Deposit(tt.deposit)
					if errors.Is(err, tt.expectedError) {
						t.Errorf("expected %v instead of %v", tt.expectedError, err)
					}
					defer wg.Done()
				}()
			}
			wg.Wait()
			if tt.wallet.balance != tt.want {
				t.Errorf("expected %v of amount instead of %v", tt.want, tt.wallet.balance)
			}
		})
	}
}

func TestWallet_DepositAndWithdrawConcurrent(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          Bitcoin
		withdraw      Bitcoin
		deposit       Bitcoin
		name          string
		expectedError error
	}{
		{
			wallet:        New("Nikita", 5.00),
			want:          8.00,
			withdraw:      1.00,
			deposit:       2.00,
			expectedError: ErrIncorrectInput,
			name:          "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			for i := 0; i < 3; i++ {
				wg.Add(2)
				go func() {
					defer wg.Done()
					err := tt.wallet.Withdraw(tt.withdraw)
					if errors.Is(err, tt.expectedError) {
						t.Errorf("expected %v instead of %v", tt.expectedError, err)
					}
				}()
				go func() {
					defer wg.Done()
					err := tt.wallet.Deposit(tt.deposit)
					if err != nil {
						t.Errorf("no error expected but got %v", err)
					}
				}()
			}
			wg.Wait()
			if tt.wallet.balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.balance)
			}
		})
	}
}

func ExampleWallet_Deposit() {
	wallet := New("Nikita", 0.00)
	_ = wallet.Deposit(2)
	fmt.Println(wallet)
	// Output:
	// User "Nikita" balance is 2.0000 BTC
}

func ExampleWallet_Withdraw() {
	wallet := New("Nikita", 2.00)
	_ = wallet.Withdraw(1)
	fmt.Println(wallet)
	// Output:
	// User "Nikita" balance is 1.0000 BTC
}
