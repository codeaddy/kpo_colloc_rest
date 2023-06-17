package product_order

import (
	"context"
)

type ProductOrderRepository interface {
	Create(ctx context.Context, productOrder ProductOrder) (int, error)
	GetById(ctx context.Context, productOrderId int) (ProductOrder, error)
	GetAll(ctx context.Context) ([]*ProductOrder, error)
}

type Service struct {
	repository ProductOrderRepository
}

func NewService(repo ProductOrderRepository) *Service {
	return &Service{repository: repo}
}

func (s *Service) Create(ctx context.Context, productOrder ProductOrder) (int, error) {
	id, err := s.repository.Create(ctx, productOrder)
	return id, err
}

func (s *Service) GetById(ctx context.Context, productOrderId int) (ProductOrder, error) {
	product, err := s.repository.GetById(ctx, productOrderId)
	if err != nil {
		return ProductOrder{}, err
	}
	return product, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*ProductOrder, error) {
	product, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}
