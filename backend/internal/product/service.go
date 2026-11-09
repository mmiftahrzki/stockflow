package product

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type stockMovement struct{}
type Service interface {
	List(ctx context.Context, query QueryList) ([]modelProduct, error)
	Create(ctx context.Context, command commandCreate) error
	Update(ctx context.Context, command commandUpdate) error
	Delete(ctx context.Context, command commandDelete) error
	CreateStock(ctx context.Context, command commandCreateStock) error
	GetMovements(ctx context.Context, query QueryMovements) ([]stockMovement, error)
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

func (svc *service) Create(ctx context.Context, command commandCreate) error {
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

func (svc *service) List(ctx context.Context, query QueryList) ([]modelProduct, error) {
	products := []modelProduct{}

	productSqls, err := svc.repo.Select(ctx, query.search, query.page, query.limit)
	if err != nil {
		return nil, err
	}

	for _, productSql := range productSqls {
		customer := modelProduct{
			Id:        int(productSql.id.Int16),
			SKU:       productSql.sku.String,
			Name:      productSql.name.String,
			Price:     int(productSql.price.Int16),
			Stock:     int(productSql.stock.Int16),
			IsActive:  productSql.is_active.Bool,
			CreatedAt: productSql.created_at.Time,
		}

		products = append(products, customer)
	}

	return products, nil
}

func (svc *service) Update(ctx context.Context, command commandUpdate) error { panic(("implement me")) }
func (svc *service) Delete(ctx context.Context, command commandDelete) error { panic(("implement me")) }
func (svc *service) CreateStock(ctx context.Context, command commandCreateStock) error {
	panic(("implement me"))
}
func (svc *service) GetMovements(ctx context.Context, query QueryMovements) ([]stockMovement, error) {
	panic(("implement me"))
}
