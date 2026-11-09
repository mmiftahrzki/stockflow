package product

import (
	"backend/internal/logger"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const LIMIT = 100
const MAX_LIMIT = 1000

type Repo interface {
	Insert(ctx context.Context, payload commandCreate) error
	Select(ctx context.Context, search string, page int, limit int) ([]modelSQL, error)
}

type repo struct {
	db  *pgxpool.Pool
	log *logrus.Entry
}

func mapSqlRowsIntoModelSql(rows pgx.Rows) ([]modelSQL, error) {
	var modelSql modelSQL
	var modelSqls []modelSQL
	var err error

	for rows.Next() {
		err = rows.Scan(
			&modelSql.id,
			&modelSql.sku,
			&modelSql.name,
			&modelSql.price,
			&modelSql.stock,
			&modelSql.is_active,
			&modelSql.created_at,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		modelSqls = append(modelSqls, modelSql)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("row iteration failed: %w", err)
	}

	return modelSqls, nil
}

func newRepo(db *pgxpool.Pool) Repo {
	return &repo{
		db:  db,
		log: logger.GetLogger().WithField("component", "customerRepo"),
	}
}

func (r *repo) Select(ctx context.Context, search string, page int, limit int) ([]modelSQL, error) {
	var products []modelSQL
	var sqlQuery string

	if limit < LIMIT {
		limit = LIMIT
	}

	if limit > MAX_LIMIT {
		limit = MAX_LIMIT
	}

	offset := page * limit

	err := ctx.Err()
	if err != nil {
		return nil, err
	}

	sqlQuery = `
		SELECT p.id,
			p.sku,
			p.name,
			p.price,
			p.stock,
			p.is_active,
			p.created_at
		FROM products p`

	if search == "" {
		sqlQuery += " WHERE p.is_active"
	} else {
		sqlQuery += " WHERE p.is_active = TRUE AND (p.name = $1 OR p.sku = $2)"
	}

	sqlQuery += fmt.Sprintf(" ORDER BY name ASC LIMIT %d OFFSET %d;", limit, offset)

	rows, err := r.db.Query(ctx, sqlQuery, search, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err = mapSqlRowsIntoModelSql(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repo) Insert(ctx context.Context, command commandCreate) error {
	err := ctx.Err()
	if err != nil {
		return err
	}

	sqlQuery := `INSERT INTO products(sku, name, price)VALUES ($1, $2, $3);`

	_, err = r.db.Exec(ctx, sqlQuery, command.SKU, command.Name, command.Price)
	if err != nil {
		return err
	}

	return nil
}
