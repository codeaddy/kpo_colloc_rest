package postgresql

import (
	"colloc_rest/internal/app/pkg/db"
	"colloc_rest/internal/app/pkg/product_order"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type ProductOrderRepo struct {
	db db.DBops
}

func NewProductOrder(db db.DBops) *ProductOrderRepo {
	return &ProductOrderRepo{db: db}
}

func (r *ProductOrderRepo) Create(ctx context.Context, product_order product_order.ProductOrder) (int, error) {
	var id int
	err := r.db.ExecQueryRow(ctx, "INSERT INTO public.product_order(product_id, order_id, quantity) VALUES ($1, $2, $3) RETURNING ID",
		product_order.ProductID, product_order.OrderID, product_order.Quantity).Scan(&id)
	return id, err
}

func (r *ProductOrderRepo) GetById(ctx context.Context, id int) (product_order.ProductOrder, error) {
	var p product_order.ProductOrder
	err := r.db.Get(ctx, &p, "SELECT id,product_id, order_id, quantity,created_at FROM public.product_order WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return product_order.ProductOrder{}, ErrObjectNotFound
	}
	return p, err
}

func (r *ProductOrderRepo) GetAll(ctx context.Context) ([]*product_order.ProductOrder, error) {
	p := make([]*product_order.ProductOrder, 0)
	err := r.db.Select(ctx, &p, "SELECT id,product_id, order_id, quantity,created_at FROM public.product_order")
	return p, err
}
