package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vo1dFl0w/test_api/internal/app/model"
)

func (s *Server) GetWallet(uuid string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			s.error(w, r, http.StatusMethodNotAllowed, ErrMethodNotAllowed)
			return
		}

		wallet, err := s.store.Wallet().GetWallet(uuid)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("an error occured: %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, wallet)
	}
}

func (s *Server) Transaction() http.HandlerFunc {
	type request struct {
		UUID      string  `json:"uuid"`
		Amount    float64 `json:"amount"`
		Operation string  `json:"operation"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.error(w, r, http.StatusMethodNotAllowed, ErrMethodNotAllowed)
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		walletID := &model.Wallet{
			UUID: req.UUID,
		}

		res, err := s.store.Wallet().Transaction(walletID, req.Amount, req.Operation)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, res)
	}
}