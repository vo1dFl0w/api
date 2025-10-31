package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vo1dFl0w/test_api/internal/app/model"
	"github.com/vo1dFl0w/test_api/internal/app/store/repository"
)

func TestWalletRepository_GetWallet(t *testing.T) {
	db, close := repository.TestDB(t, databaseURL)
	defer close()

	s := repository.New(db)
	w1 := model.TestWallet(t)

	res1, err1 := s.Wallet().GetWallet(w1.UUID)

	assert.NoError(t, err1)
	assert.NotNil(t, res1)

	w2 := &model.Wallet{UUID: "invalid"}
	_, err2 := s.Wallet().GetWallet(w2.UUID)
	assert.Error(t, err2)
}

func TestWalletRepository_Transaction(t *testing.T) {
	db, close := repository.TestDB(t, databaseURL)
	defer close()

	s := repository.New(db)
	w := model.TestWallet(t)

	testCases := []struct {
		name       string
		wallet     model.Wallet
		operation  string
		amount     float64
		expAccount float64
		expError   bool
	}{
		{
			name:       "valid deposit",
			wallet:     model.Wallet{UUID: w.UUID, Account: w.Account},
			operation:  "DEPOSIT",
			amount:     1000.0,
			expAccount: w.Account + 1000.0,
			expError:   false,
		},
		{
			name:       "valid withdraw",
			wallet:     model.Wallet{UUID: w.UUID, Account: w.Account},
			operation:  "WITHDRAW",
			amount:     1000.0,
			expAccount: w.Account - 1000.0,
			expError:   false,
		},
		{
			name:       "invalid operation",
			wallet:     model.Wallet{UUID: w.UUID, Account: 0.0},
			operation:  "",
			amount:     1000.0,
			expAccount: 0.0,
			expError:   true,
		},
		{
			name:       "invalid deposit",
			wallet:     model.Wallet{UUID: "", Account: 0.0},
			operation:  "DEPOSIT",
			amount:     1000.0,
			expAccount: 0.0,
			expError:   true,
		},
		{
			name:       "invalid withdraw",
			wallet:     model.Wallet{UUID: w.UUID, Account: 0.0},
			operation:  "DEPOSIT",
			amount:     1000.0,
			expAccount: w.Account - 1000.0,
			expError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := s.Wallet().Transaction(&tc.wallet, tc.amount, tc.operation)

			if tc.expError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
				assert.Equal(t, tc.expAccount, res.Account)
			}
		})
	}
}
