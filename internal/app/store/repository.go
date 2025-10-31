package store

import "github.com/vo1dFl0w/test_api/internal/app/model"

type WalletRepository interface {
	GetWallet(uuid string) (*model.Wallet, error)
	Transaction(w *model.Wallet, sum float64, operation string) (*model.Wallet, error)
}
