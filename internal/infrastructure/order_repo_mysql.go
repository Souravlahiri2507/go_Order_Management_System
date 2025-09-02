package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"go_oms/internal/domain"
)

type MySQLOrderRepository struct {
	db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) domain.OrderRepository {
	return &MySQLOrderRepository{db: db}
}

func (r *MySQLOrderRepository) Create(ctx context.Context, o *domain.Order) (int64, error) {
	q := `INSERT INTO orders(customer_name, item, quantity, price, status) VALUES (?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, q, o.CustomerName, o.Item, o.Quantity, o.Price, o.Status)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *MySQLOrderRepository) GetByID(ctx context.Context, id int64) (*domain.Order, error) {
	q := `SELECT id, customer_name, item, quantity, price, status, created_at, updated_at FROM orders WHERE id = ?`
	row := r.db.QueryRowContext(ctx, q, id)
	o := &domain.Order{}
	if err := row.Scan(&o.ID, &o.CustomerName, &o.Item, &o.Quantity, &o.Price, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return o, nil
}

func (r *MySQLOrderRepository) Update(ctx context.Context, o *domain.Order) error {
	q := `UPDATE orders SET customer_name=?, item=?, quantity=?, price=?, status=?, updated_at=? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, q, o.CustomerName, o.Item, o.Quantity, o.Price, o.Status, time.Now(), o.ID)
	return err
}

func (r *MySQLOrderRepository) Delete(ctx context.Context, id int64) error {
	q := `DELETE FROM orders WHERE id = ?`
	_, err := r.db.ExecContext(ctx, q, id)
	return err
}

func (r *MySQLOrderRepository) List(ctx context.Context) ([]*domain.Order, error) {
	q := `SELECT id, customer_name, item, quantity, price, status, created_at, updated_at FROM orders ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*domain.Order
	for rows.Next() {
		o := &domain.Order{}
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.Item, &o.Quantity, &o.Price, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return out, rows.Err()
}
