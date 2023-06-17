package product_order

import "time"

type ProductOrder struct {
	ID        int       `db:"id"`
	ProductID int       `db:"product_id"`
	OrderID   int       `db:"order_id"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
}
