package order

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type order struct{}

func New(db *pgxpool.Pool) order {
	return order{}
}

func (p *order) Post(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (p *order) PostCancel(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (p *order) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}
