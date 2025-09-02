package domain

import "context"

type OrderRepository interface {
Create(ctx context.Context, o *Order) (int64, error)
GetByID(ctx context.Context, id int64) (*Order, error)
Update(ctx context.Context, o *Order) error
Delete(ctx context.Context, id int64) error
List(ctx context.Context) ([]*Order, error)
}
