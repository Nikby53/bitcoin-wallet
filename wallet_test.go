package main

import (
	"errors"
	"sync"
	"testing"
)

func TestWallet_Deposit(t *testing.T) {
	tests := []struct {
		wallet  *Wallet
		want    float64
		deposit float64
		name    string
	}{
		{
			wallet:  New("Nikita", 1.00),
			want:    5.32,
			deposit: 4.32,
			name:    "success",
		},
		{
			wallet:  New("Nikita", 1.00),
			want:    1.00,
			deposit: 0.00,
			name:    "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wallet.Deposit(tt.deposit)
			if tt.wallet.Balance != Bitcoin(tt.want) {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.Balance)
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          float64
		withdraw      float64
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
			expectedError: errors.New("not enough money to withdraw"),
			withdraw:      2,
			name:          "fail:not enough money to withdraw",
			want:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wallet.Withdraw(tt.withdraw)
			if tt.expectedError != nil && tt.expectedError.Error() != err.Error() {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
			if tt.wallet.Balance != Bitcoin(tt.want) {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.Balance)
			}
		})
	}
}

func TestWallet_WithdrawConcurrent(t *testing.T) {
	tests := []struct {
		wallet        *Wallet
		want          float64
		withdraw      float64
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
			expectedError: errors.New("not enough money to withdraw"),
			withdraw:      2.00,
			name:          "fail:not enough money to withdraw",
			want:          1.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(3)
			for i := 0; i < 3; i++ {
				go func() {
					err := tt.wallet.Withdraw(tt.withdraw)
					if tt.expectedError != nil && tt.expectedError.Error() != err.Error() {
						t.Errorf("expected %v instead of %v", tt.expectedError, err)
					}
					defer wg.Done()
				}()
			}
			wg.Wait()
			if tt.wallet.Balance != Bitcoin(tt.want) {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.Balance)
			}
		})
	}
}

func TestWallet_DepositConcurrent(t *testing.T) {
	tests := []struct {
		wallet  *Wallet
		want    float64
		deposit float64
		name    string
	}{
		{
			wallet:  New("Nikita", 1.00),
			want:    7.00,
			deposit: 2.00,
			name:    "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(3)
			for i := 0; i < 3; i++ {
				go func() {
					tt.wallet.Deposit(tt.deposit)
					defer wg.Done()
				}()
			}
			wg.Wait()
			if tt.wallet.Balance != Bitcoin(tt.want) {
				t.Errorf("expected %v of amount instead of %v", tt.want, tt.wallet.Balance)
			}
		})
	}
}

func ExampleWallet_Deposit() {
	tests := []struct {
		wallet  *Wallet
		want    float64
		deposit float64
	}{
		{
			wallet:  New("Nikita", 1.00),
			want:    5.32,
			deposit: 4.32,
		},
	}
	for _, tt := range tests {
		tt.wallet.Deposit(tt.deposit)
		// Output:
		// 5.32
	}
}

func ExampleWallet_Withdraw() {
	tests := []struct {
		wallet   *Wallet
		want     float64
		withdraw float64
	}{
		{
			wallet:   New("Nikita", 1.00),
			want:     0.50,
			withdraw: 0.50,
		},
	}
	for _, tt := range tests {
		tt.wallet.Deposit(tt.withdraw)
		// Output:
		// 0.50
	}
}
