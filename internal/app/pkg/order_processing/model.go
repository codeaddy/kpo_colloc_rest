package order_processing

import "colloc_rest/internal/app/pkg/product"

type Order struct {
	ID         int               `db:"id"`
	UserID     int               `db:"user_id"`
	Products   []product.Product `db:"products"`
	Status     string            `db:"status"`
	TotalPrice int64             `db:"total_price"`
}
