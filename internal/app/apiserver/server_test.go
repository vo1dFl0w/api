package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vo1dFl0w/test_api/internal/app/model"
	"github.com/vo1dFl0w/test_api/internal/app/store/testrepository"
)

func TestServer_GetWallet(t *testing.T) {
	s := NewServer(testrepository.New())
	w := model.TestWallet(t)

	testCases := []struct {
		name    string
		uuid    string
		expCode int
	}{
		{
			name:    "valid",
			uuid:    w.UUID,
			expCode: http.StatusOK,
		},
		{
			name:    "invalid",
			uuid:    "",
			expCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/api/v1/wallets/%s", tc.uuid)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expCode, rec.Code)
		})
	}
}

func TestServer_Transaction(t *testing.T) {
	s := NewServer(testrepository.New())
	w := model.TestWallet(t)

	testCases := []struct {
		name       string
		payload    interface{}
		expAccount float64
		expCode    int
	}{
		{
			name: "valid deposit",
			payload: map[string]interface{}{
				"uuid":      w.UUID,
				"operation": "DEPOSIT",
				"amount":    1000.00,
			},
			expCode: http.StatusOK,
		},
		{
			name: "valid withdraw",
			payload: map[string]interface{}{
				"uuid":      w.UUID,
				"operation": "WITHDRAW",
				"amount":    1000.00,
			},
			expCode: http.StatusOK,
		},
		{
			name: "invalid uuid",
			payload: map[string]interface{}{
				"uuid": "",
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "invalid operation",
			payload: map[string]interface{}{
				"uuid":      w.UUID,
				"operation": "",
				"amount":    1000.00,
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "invalid withdraw",
			payload: map[string]interface{}{
				"uuid":      w.UUID,
				"operation": "WITHDRAW",
				"amount":    1000000000.00,
			},
			expCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expCode, rec.Code)
		})
	}
}
