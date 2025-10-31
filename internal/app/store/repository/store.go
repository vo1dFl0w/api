package repository

import (
	"database/sql"

	"github.com/vo1dFl0w/test_api/internal/app/store"
)

type Store struct {
	db *sql.DB
	walletRepository *WalletRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Wallet() store.WalletRepository {
	if s.walletRepository != nil {
		return s.walletRepository
	}

	s.walletRepository = &WalletRepository{
		store: s,
	}

	return s.walletRepository
}