package testrepository

import (
	"github.com/vo1dFl0w/test_api/internal/app/store"
)

type Store struct {
	walletRepository *WalletRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Wallet() store.WalletRepository {
	if s.walletRepository != nil {
		return s.walletRepository
	}

	s.walletRepository = &WalletRepository{
		store: s,
		wallets: map[string]float64{"f81d4fae-7dec-11d0-a765-00a0c91e6bf6": 10000.00},
	}

	return s.walletRepository
}