package apiserver

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/vo1dFl0w/test_api/internal/app/apiserver/config"
	"github.com/vo1dFl0w/test_api/internal/app/store"
	"github.com/vo1dFl0w/test_api/internal/app/store/repository"
)

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
)

type Server struct {
	store  store.Store
	router *http.ServeMux
}

func NewServer(store store.Store) *Server {
	s := &Server{
		store:  store,
		router: http.NewServeMux(),
	}
	s.configureRouter()

	return s
}

func Run(cfg *config.Config) error {
	db, err := initDB(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to init database: %w", err)
	}
	defer db.Close()

	store := repository.New(db)

	srv := NewServer(store)

	s := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: srv,
	}
	log.Printf("server has started %s", cfg.HTTPAddr)
	return http.ListenAndServe(s.Addr, s.Handler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/api/v1/wallet", s.Transaction())

	s.router.HandleFunc("/api/v1/wallets/", func(w http.ResponseWriter, r *http.Request) {
		uuid := strings.TrimPrefix(r.URL.Path, "/api/v1/wallets/")

		s.GetWallet(uuid).ServeHTTP(w, r)
	})
}

func initDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(300)
	db.SetMaxIdleConns(150)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
