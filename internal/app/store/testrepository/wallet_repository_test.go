package testrepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vo1dFl0w/test_api/internal/app/model"
	"github.com/vo1dFl0w/test_api/internal/app/store/testrepository"
)

func Test_WalletRepository_GetWallet(t *testing.T) {
	s := testrepository.New()
	w := model.TestWallet(t)

	w, err := s.Wallet().GetWallet(w.UUID)
	assert.NoError(t, err)
	assert.NotNil(t, w)

	w.UUID = ""

	w, err = s.Wallet().GetWallet(w.UUID)
	assert.Error(t, err)
}

func Test_WalletRepository_Transaction(t *testing.T) {
	s := testrepository.New()
	w := model.TestWallet(t)

	testCases := []struct {
		name       string
		wallet     model.Wallet
		amount     float64
		operation  string
		expAccount float64
		expErr     bool
	}{
		{
			name:       "valid_deposit",
			wallet:     model.Wallet{UUID: w.UUID, Account: w.Account},
			amount:     1000.00,
			operation:  "DEPOSIT",
			expAccount: w.Account + 1000.00,
			expErr:     false,
		},
		{
			name:       "valid_withdraw",
			wallet:     model.Wallet{UUID: w.UUID, Account: w.Account},
			amount:     1000.00,
			operation:  "WITHDRAW",
			expAccount: w.Account - 1000.00,
			expErr:     false,
		},
		{
			name:       "invalid_operation",
			wallet:     model.Wallet{UUID: w.UUID, Account: w.Account},
			amount:     1000.00,
			operation:  "invalid",
			expAccount: w.Account,
			expErr:     true,
		},
		{
			name:       "invalid_id",
			wallet:     model.Wallet{UUID: "", Account: 0.0},
			amount:     1000.00,
			operation:  "invalid_id",
			expAccount: 0.0,
			expErr:     true,
		},
		{
			name:       "invalid_withdraw",
			wallet:     model.Wallet{UUID: w.UUID, Account: w.Account},
			amount:     1000000.00,
			operation:  "WITHDRAW",
			expAccount: w.Account,
			expErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := s.Wallet().Transaction(&tc.wallet, tc.amount, tc.operation)
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
				assert.Equal(t, tc.expAccount, res.Account)
			}
		})
	}
}
