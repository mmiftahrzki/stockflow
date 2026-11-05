package product

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateProduct(ctx context.Context, command createCommand) error
}

type service struct {
	repo Repo
	log  *logrus.Entry
}

func newService(r Repo) Service {
	svc := &service{
		repo: r,
	}

	return svc
}

func (svc *service) CreateProduct(ctx context.Context, command createCommand) error {
	err := ctx.Err()
	if err != nil {
		return err
	}

	err = svc.repo.Insert(ctx, command)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errProductSKUAlreadyExists
			}
		}

		return err
	}

	return nil
}
