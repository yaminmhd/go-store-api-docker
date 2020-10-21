package contract

import "time"

type GetProducts struct {
	Products   []ProductSummary `json:"products"`
	TotalCount int            `json:"total_count"`
}

type ProductSummary struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  uint16    `json:"quantity"`
	State     string    `json:"state"`
}
