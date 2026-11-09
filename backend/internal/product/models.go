package product

import (
	"database/sql"
	"time"
)

type modelSQL struct {
	id         sql.NullInt16
	sku        sql.NullString
	name       sql.NullString
	price      sql.NullInt16
	stock      sql.NullInt16
	is_active  sql.NullBool
	created_at sql.NullTime
}

type modelProduct struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
