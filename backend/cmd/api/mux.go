package main

import (
	"backend/internal/order"
	"backend/internal/product"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newMux(db *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	product := product.New(db)
	order := order.New(db)

	mux.HandleFunc("GET /api/products?search=&page=&limit= ", product.Get)
	mux.HandleFunc("POST /api/products", product.Post)
	mux.HandleFunc("PUT /api/products/{id}", product.Put)
	mux.HandleFunc("DELETE /api/products/{id}", product.Delete)
	mux.HandleFunc("POST /api/products/{id}/stock", product.PostStock)
	mux.HandleFunc("GET /api/products/{id}/movements", product.GetMovements)
	mux.HandleFunc("POST /api/orders", order.Post)
	mux.HandleFunc("POST /api/orders/{id}/cancel", order.PostCancel)
	mux.HandleFunc("GET /api/orders/{id}", order.Get)

	return mux
}
