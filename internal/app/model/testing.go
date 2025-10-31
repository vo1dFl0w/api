package model

import "testing"

func TestWallet(t *testing.T) *Wallet {
	return &Wallet{
		UUID: "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		Account: 10000.00,
	}
}
