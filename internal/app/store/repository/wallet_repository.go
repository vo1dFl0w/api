package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/vo1dFl0w/test_api/internal/app/model"
)

type WalletRepository struct {
	store *Store
}

func (r *WalletRepository) GetWallet(uuid string) (*model.Wallet, error) {
	w := &model.Wallet{}

	var accountDB int64
	if err := r.store.db.QueryRow(
		"SELECT uuid, account FROM wallets WHERE uuid = $1",
		uuid,
	).Scan(&w.UUID, &accountDB); err != nil {
		return nil, fmt.Errorf("cannot find wallet by uuid %s: %w", uuid, err)
	}

	w.Account = float64(accountDB) / 100.0

	return w, nil
}

func (r *WalletRepository) Transaction(w *model.Wallet, amount float64, operation string) (*model.Wallet, error) {
	const maxAttempts = 5
    baseDelay := 100 * time.Millisecond
    maxDelay := 5 * time.Second

    rand.Seed(time.Now().UnixNano())

    amountDB := int64(math.Round(amount * 100.0))
    var lastErr error

    for attempt := 1; attempt <= maxAttempts; attempt++ {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

        res, err := func() (*model.Wallet, error) {
            defer cancel()

            tx, err := r.store.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
            if err != nil {
                return nil, fmt.Errorf("begin tx: %w", err)
            }
            defer func() { _ = tx.Rollback() }()

            var accountDB int64
            op := strings.ToUpper(operation)

            switch op {
            case "DEPOSIT":
                err = tx.QueryRowContext(ctx,
                    "UPDATE wallets SET account = account + $1 WHERE uuid = $2 RETURNING account",
                    amountDB, w.UUID).Scan(&accountDB)

            case "WITHDRAW":
                err = tx.QueryRowContext(ctx,
                    "UPDATE wallets SET account = account - $1 WHERE uuid = $2 AND account >= $1 RETURNING account",
                    amountDB, w.UUID).Scan(&accountDB)

            default:
                return nil, fmt.Errorf("invalid operation: %s", operation)
            }

            if err != nil {
                if errors.Is(err, sql.ErrNoRows) {
					return nil, fmt.Errorf("insufficient funds or wallet not found")
                }
				return nil, fmt.Errorf("update failed: %w", err)
            }

            if err := tx.Commit(); err != nil {
                return nil, fmt.Errorf("commit failed: %w", err)
            }

            return &model.Wallet{UUID: w.UUID, Account: float64(accountDB) / 100.0}, nil
        }()

        if err == nil {
            return res, nil
        }

        lastErr = err

        retry := false
        var pqErr *pq.Error
        if errors.As(err, &pqErr) {
            code := string(pqErr.Code)
            if code == "40001" || code == "40P01" {
                retry = true
            }
        }
        if !retry {
            s := strings.ToLower(err.Error())
            if strings.Contains(s, "could not serialize") || strings.Contains(s, "serialization_failure") || strings.Contains(s, "deadlock") {
                retry = true
            }
        }

        if retry && attempt < maxAttempts {
            backoffTime := baseDelay * time.Duration(math.Pow(2, float64(attempt)))
            if backoffTime > maxDelay {
                backoffTime = maxDelay
            }
            jitter := time.Duration(rand.Int63n(int64(backoffTime)))
            time.Sleep(jitter)
            continue
        }

        return nil, lastErr
    }

    return nil, fmt.Errorf("transaction failed after %d attempts: %w", maxAttempts, lastErr)
}
