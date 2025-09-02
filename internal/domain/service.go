package domain

import (
"context"
"encoding/json"
"fmt"

"github.com/segmentio/kafka-go"
)

type OrderService struct {
repo        OrderRepository
kafkaWriter *kafka.Writer
}

func NewOrderService(repo OrderRepository, kafkaWriter *kafka.Writer) *OrderService {
return &OrderService{repo: repo, kafkaWriter: kafkaWriter}
}

func (s *OrderService) CreateOrder(ctx context.Context, o *Order) (int64, error) {
id, err := s.repo.Create(ctx, o)
if err != nil {
return 0, err
}
o.ID = id
_ = s.publish(ctx, "order_created", o.ID, o)
return id, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id int64) (*Order, error) {
return s.repo.GetByID(ctx, id)
}

func (s *OrderService) UpdateOrder(ctx context.Context, o *Order) error {
if err := s.repo.Update(ctx, o); err != nil {
return err
}
_ = s.publish(ctx, "order_updated", o.ID, o)
return nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, id int64) error {
if err := s.repo.Delete(ctx, id); err != nil {
return err
}
_ = s.publish(ctx, "order_deleted", id, map[string]any{"deleted": true})
return nil
}

func (s *OrderService) ListOrders(ctx context.Context) ([]*Order, error) {
return s.repo.List(ctx)
}

func (s *OrderService) publish(ctx context.Context, event string, orderID int64, payload any) error {
ev := map[string]any{
"event_type": event,
"order_id":   orderID,
"payload":    payload,
}
b, _ := json.Marshal(ev)
return s.kafkaWriter.WriteMessages(ctx, kafka.Message{
Key:   []byte(fmt.Sprintf("%d", orderID)),
Value: b,
})
}
