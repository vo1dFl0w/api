package store

type Store interface {
	Wallet() WalletRepository
}