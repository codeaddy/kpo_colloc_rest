package product

import "time"

type Product struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	Quantity    int       `db:"quantity"`
	CreatedAt   time.Time `db:"created_at"`
}
