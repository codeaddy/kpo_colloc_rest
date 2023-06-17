package postgresql

import (
	"colloc_rest/internal/app/pkg/db"
	"colloc_rest/internal/app/pkg/order"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type OrderRepo struct {
	db db.DBops
}

func NewOrder(db db.DBops) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(ctx context.Context, orderRow order.OrderRow) (int, error) {
	var id int
	err := r.db.ExecQueryRow(ctx, "INSERT INTO public.order(user_id, status) VALUES ($1, $2) RETURNING ID",
		orderRow.UserID, orderRow.Status).Scan(&id)
	return id, err
}

func (r *OrderRepo) GetById(ctx context.Context, id int) (order.OrderRow, error) {
	var o order.OrderRow
	err := r.db.Get(ctx, &o, "SELECT id,user_id,status,created_at FROM public.order WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return order.OrderRow{}, ErrObjectNotFound
	}
	return o, err
}

func (r *OrderRepo) GetAll(ctx context.Context) ([]*order.OrderRow, error) {
	o := make([]*order.OrderRow, 0)
	err := r.db.Select(ctx, &o, "SELECT id,user_id,status,created_at FROM public.order")
	return o, err
}

func (r *OrderRepo) GetAllByUserId(ctx context.Context, userID int) ([]*order.OrderRow, error) {
	o := make([]*order.OrderRow, 0)
	err := r.db.Select(ctx, &o, "SELECT id,user_id,status,created_at FROM public.order where user_id=$1", userID)
	if err == sql.ErrNoRows {
		return nil, ErrObjectNotFound
	}
	return o, err
}
