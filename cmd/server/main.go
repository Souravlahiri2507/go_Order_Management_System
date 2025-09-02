package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "go_oms/internal/api"
	// "go_oms/internal/domain"
	"go_oms/internal/infrastructure"
)

func main() {
	// load .env if present
	_ = godotenv.Load()

	// config
	cfg := infrastructure.NewConfigFromEnv()

	// DB
	db, err := infrastructure.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	// Kafka (producer + consumer)
	writer := infrastructure.NewKafkaWriter(cfg)
	defer writer.Close()
	reader := infrastructure.NewKafkaReader(cfg, "order-events-consumer-group")

	// repository + service
	//repo := infrastructure.NewMySQLOrderRepository(db)
	//svc := domain.NewOrderService(repo, writer)

	// start consumer goroutine (logs and stores events)
	go infrastructure.StartOrderEventsConsumer(context.Background(), reader, db)

	// HTTP server
	r := gin.Default()
	//api.RegisterOrderRoutes(r, svc)

	log.Printf("server listening on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
