package order

import (
	"backend/internal/product"
	"context"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Create(ctx context.Context, command commandCreate) error
	Cancel(ctx context.Context, command commandCancel) error
	Detail(ctx context.Context, query queryDetail) error
}

type service struct {
	log  *logrus.Entry
	repo Repo
}

var _productService product.Service = nil

func newService(r Repo, productService product.Service) Service {
	_productService = productService

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

	// skus := []string{}
	// for i := 0; i < len(command.Items); i++ {
	// 	skus = append(skus, command.Items[i].SKU)
	// }

	// queryProduct := product.QueryList{
	// 	SKUs: skus,
	// }
	// products, err := _productService.List(ctx, queryProduct)
	// if err != nil {
	// 	panic("implement me!")
	// }

	// err = svc.repo.Insert(ctx, command)
	// if err != nil {
	// 	var pgErr *pgconn.PgError

	// 	if errors.As(err, &pgErr) {
	// 	}

	// 	return err
	// }

	return nil
}

func (svc *service) Cancel(ctx context.Context, command commandCancel) error {
	panic("implement me!")
}

func (svc *service) Detail(ctx context.Context, query queryDetail) error {
	panic("implement me!")
}
