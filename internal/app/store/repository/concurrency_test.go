package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vo1dFl0w/test_api/internal/app/model"
	"github.com/vo1dFl0w/test_api/internal/app/store/repository"
)

func TestConcurrency(t *testing.T) {
	db, close := repository.TestDB(t, databaseURL)
	defer close()

	s := repository.New(db)
	w := model.TestWallet(t)

	req := 1000
	amount := 1.0
	operation := "DEPOSIT"
	errCount := 0

	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < req; i++ {
		go func() {
			defer wg.Done()
			_, err := s.Wallet().Transaction(w, amount, operation)
			if err != nil {
				errCount++
			}
		}()
	}

	res, err := s.Wallet().GetWallet(w.UUID)
	if err != nil {
		t.Fatal(err)
	}

	expAccount := float64(req) + w.Account
	
	assert.Equal(t, expAccount, res.Account)

	wg.Wait()
}
