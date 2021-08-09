package wallet

import (
	"sync"
	"testing"
)

func TestWallet_Deposit(t *testing.T) {
	tests := []struct {
		wallet  *Wallet
		want    Bitcoin
		deposit Bitcoin
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
			if tt.wallet.Balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.Balance)
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
			expectedError: AppError,
			withdraw:      2,
			name:          "fail:not enough money to withdraw",
			want:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wallet.Withdraw(tt.withdraw)
			if tt.expectedError != err {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
			if tt.wallet.Balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.Balance)
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
			expectedError: AppError,
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
					if tt.expectedError != err {
						t.Errorf("expected %v instead of %v", tt.expectedError, err)
					}
					defer wg.Done()
				}()
			}
			wg.Wait()
			if tt.wallet.Balance != tt.want {
				t.Errorf("expected %v of amount intead of %v", tt.want, tt.wallet.Balance)
			}
		})
	}
}

func TestWallet_DepositConcurrent(t *testing.T) {
	tests := []struct {
		wallet  *Wallet
		want    Bitcoin
		deposit Bitcoin
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
			for i := 0; i < 3; i++ {
				wg.Add(1)
				go func() {
					tt.wallet.Deposit(tt.deposit)
					defer wg.Done()
				}()
			}
			wg.Wait()
			if tt.wallet.Balance != tt.want {
				t.Errorf("expected %v of amount instead of %v", tt.want, tt.wallet.Balance)
			}
		})
	}
}

func ExampleWallet_Deposit() {
	tests := []struct {
		wallet  *Wallet
		want    Bitcoin
		deposit Bitcoin
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
		want     Bitcoin
		withdraw Bitcoin
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
