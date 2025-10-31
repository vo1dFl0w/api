package testrepository

import (
	"fmt"
	"strings"

	"github.com/vo1dFl0w/test_api/internal/app/model"
)

type WalletRepository struct {
	store   *Store
	wallets map[string]float64
}

func (r *WalletRepository) GetWallet(uuid string) (*model.Wallet, error) {

	_, ok := r.wallets[uuid]
	if !ok {
		return nil, fmt.Errorf("cannot find wallet by uuid %s", uuid)
	}

	return &model.Wallet{UUID: uuid, Account: r.wallets[uuid]}, nil
}

func (r *WalletRepository) Transaction(w *model.Wallet, amount float64, operation string) (*model.Wallet, error) {
	account, ok := r.wallets[w.UUID]
	if !ok {
		return nil, fmt.Errorf("cannot find wallet by uuid %s", w.UUID)
	}

	switch strings.ToUpper(operation) {
	case "DEPOSIT":
		account += amount
	case "WITHDRAW":
		if account < amount {
			return nil, fmt.Errorf("insufficient funds in the account")
		}
		account -= amount
	default:
		return nil, fmt.Errorf("invalid operation type: %s", operation)
	}

	return &model.Wallet{UUID: w.UUID, Account: account}, nil
}
