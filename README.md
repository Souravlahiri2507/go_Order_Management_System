# go_Order_Management_System
Run MySQL & Kafka with Docker, then `go run ./cmd/server`.
Endpoints:
- POST /orders
- GET /orders
- GET /orders/:id
- PUT /orders/:id
- DELETE /orders/:id
Events are published to Kafka and stored in `order_events`.
