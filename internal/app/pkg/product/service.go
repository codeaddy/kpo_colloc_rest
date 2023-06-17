package product

import (
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, product Product) (int, error)
	GetById(ctx context.Context, productID int) (Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
}

type Service struct {
	repository ProductRepository
}

func NewService(repo ProductRepository) *Service {
	return &Service{repository: repo}
}

func (s *Service) Create(ctx context.Context, product Product) (int, error) {
	id, err := s.repository.Create(ctx, product)
	return id, err
}

func (s *Service) GetById(ctx context.Context, productID int) (Product, error) {
	product, err := s.repository.GetById(ctx, productID)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*Product, error) {
	product, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}
