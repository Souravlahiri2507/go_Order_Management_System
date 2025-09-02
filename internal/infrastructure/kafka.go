package infrastructure

import (
"context"
"database/sql"
"encoding/json"
"log"

"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(cfg *Config) *kafka.Writer {
return kafka.NewWriter(kafka.WriterConfig{
Brokers: []string{cfg.KafkaBrokers},
Topic:   cfg.KafkaTopic,
})
}

func NewKafkaReader(cfg *Config, groupID string) *kafka.Reader {
return kafka.NewReader(kafka.ReaderConfig{
Brokers: []string{cfg.KafkaBrokers},
GroupID: groupID,
Topic:   cfg.KafkaTopic,
})
}

// Consumes order events and writes them to order_events table
func StartOrderEventsConsumer(ctx context.Context, reader *kafka.Reader, db *sql.DB) {
log.Println("Kafka consumer started...")
for {
m, err := reader.ReadMessage(ctx)
if err != nil {
log.Printf("kafka read error: %v", err)
return
}
log.Printf("event received: key=%s value=%s", string(m.Key), string(m.Value))

var payload map[string]interface{}
if err := json.Unmarshal(m.Value, &payload); err != nil {
log.Printf("json unmarshal error: %v", err)
continue
}

q := "INSERT INTO order_events(order_id, event_type, payload) VALUES (?, ?, ?)"
_, err = db.Exec(q, payload["order_id"], payload["event_type"], string(m.Value))
if err != nil {
log.Printf("db insert event error: %v", err)
continue
}
}
}
