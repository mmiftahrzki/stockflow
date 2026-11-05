package main

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newMux(db *pgxpool.Pool, appCfg AppConfig) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/products?search=&page=&limit= ", nil)
	mux.HandleFunc("POST /api/products", nil)
	mux.HandleFunc("PUT /api/products/{id}", nil)
	mux.HandleFunc("DELETE /api/products/{id}", nil)
	mux.HandleFunc("POST /api/products/{id}/stock", nil)
	mux.HandleFunc("GET /api/products/{id}/movements", nil)
	mux.HandleFunc("POST /api/orders", nil)
	mux.HandleFunc("POST /api/orders/{id}/cancel", nil)
	mux.HandleFunc("GET /api/orders/{id}", nil)

	return mux
}
