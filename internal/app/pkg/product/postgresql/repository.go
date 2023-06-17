package postgresql

import (
	"colloc_rest/internal/app/pkg/db"
	"colloc_rest/internal/app/pkg/product"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type ProductRepo struct {
	db db.DBops
}

func NewProduct(db db.DBops) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(ctx context.Context, product product.Product) (int, error) {
	var id int
	err := r.db.ExecQueryRow(ctx, "INSERT INTO public.product(name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING ID",
		product.Name, product.Description, product.Price, product.Quantity).Scan(&id)
	return id, err
}

func (r *ProductRepo) GetById(ctx context.Context, id int) (product.Product, error) {
	var p product.Product
	err := r.db.Get(ctx, &p, "SELECT id,name,description,price,quantity,created_at FROM public.product WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return product.Product{}, ErrObjectNotFound
	}
	return p, err
}

func (r *ProductRepo) GetAll(ctx context.Context) ([]*product.Product, error) {
	p := make([]*product.Product, 0)
	err := r.db.Select(ctx, &p, "SELECT id,name,description,price,quantity,created_at FROM public.product")
	return p, err
}
