package domain

import "time"

type Order struct {
ID           int64     `json:"id"`
CustomerName string    `json:"customer_name"`
Item         string    `json:"item"`
Quantity     int       `json:"quantity"`
Price        float64   `json:"price"`
Status       string    `json:"status"`
CreatedAt    time.Time `json:"created_at"`
UpdatedAt    time.Time `json:"updated_at"`
}
