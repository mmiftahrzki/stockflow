package product

import (
	"backend/internal/logger"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Repo interface {
	Insert(ctx context.Context, payload createCommand) error
}

type repo struct {
	db  *pgxpool.Pool
	log *logrus.Entry
}

func newRepo(db *pgxpool.Pool) Repo {
	return &repo{
		db:  db,
		log: logger.GetLogger().WithField("component", "customerRepo"),
	}
}

func (r *repo) Insert(ctx context.Context, command createCommand) error {
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
