package order

import (
	"time"
)

type OrderRow struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
