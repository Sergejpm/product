package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Id          string          `db:"id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Price       decimal.Decimal `db:"price"`
	CreatedAt   time.Time       `db:"created_at"`
}
