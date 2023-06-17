package order

import (
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, orderRow OrderRow) (int, error)
	GetById(ctx context.Context, orderID int) (OrderRow, error)
	GetAll(ctx context.Context) ([]*OrderRow, error)
	GetAllByUserId(ctx context.Context, userID int) ([]*OrderRow, error)
}

type Service struct {
	repository OrderRepository
}

func NewService(repo OrderRepository) *Service {
	return &Service{repository: repo}
}

func (s *Service) Create(ctx context.Context, orderRow OrderRow) (int, error) {
	id, err := s.repository.Create(ctx, orderRow)
	return id, err
}

func (s *Service) GetById(ctx context.Context, orderID int) (OrderRow, error) {
	order, err := s.repository.GetById(ctx, orderID)
	if err != nil {
		return OrderRow{}, err
	}
	return order, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*OrderRow, error) {
	order, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *Service) GetAllByUserId(ctx context.Context, userID int) ([]*OrderRow, error) {
	order, err := s.repository.GetAllByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
