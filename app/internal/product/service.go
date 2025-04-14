package product

import (
	"context"
	"log/slog"

	"github.com/ncostamagna/go-monitoring/app/internal/domain"
)

type (
	Filters struct {
		Name string
	}

	Service interface {
		Store(ctx context.Context, name, description string, price float64) (*domain.Product, error)
		Get(ctx context.Context, id int) (*domain.Product, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Product, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, name, description *string, price *float64) error
		Count(ctx context.Context, filters Filters) (int, error)
	}

	service struct {
		log  *slog.Logger
		repo Repository
	}
)

// NewService is a service handler.
func NewService(l *slog.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Store(ctx context.Context, name, description string, price float64) (*domain.Product, error) {

	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	if err := s.repo.Store(ctx, product); err != nil {
		return nil, err
	}
	s.log.Info("product stored", "product", product)
	return product, nil
}

func (s service) GetAll(ctx context.Context, _ Filters, offset, limit int) ([]domain.Product, error) {
	products, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	s.log.Info("products retrieved", "products", products)
	return products, nil
}

func (s service) Get(ctx context.Context, id int) (*domain.Product, error) {
	product, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	s.log.Info("product retrieved", "product", product)
	return product, nil
}

func (s service) Delete(ctx context.Context, id int) error {

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	s.log.Info("product deleted", "id", id)
	return nil
}

func (s service) Update(ctx context.Context, id int, name, description *string, price *float64) error {
	if err := s.repo.Update(ctx, id, name, description, price); err != nil {
		return err
	}
	s.log.Info("product updated", "id", id, "name", name, "description", description, "price", price)
	return nil
}

func (s service) Count(ctx context.Context, _ Filters) (int, error) {
	return s.repo.Count(ctx)
}
