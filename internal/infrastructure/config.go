package infrastructure

import (
"fmt"
"os"
)

type Config struct {
DBHost       string
DBPort       string
DBUser       string
DBPass       string
DBName       string
KafkaBrokers string
KafkaTopic   string
ServerPort   string
}

func NewConfigFromEnv() *Config {
return &Config{
DBHost:       getEnv("DB_HOST", "127.0.0.1"),
DBPort:       getEnv("DB_PORT", "3306"),
DBUser:       getEnv("DB_USER", "root"),
DBPass:       getEnv("DB_PASS", "password"),
DBName:       getEnv("DB_NAME", "go_commerce"),
KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
KafkaTopic:   getEnv("KAFKA_TOPIC", "order-events"),
ServerPort:   getEnv("PORT", "8080"),
}
}

func getEnv(k, d string) string {
if v := os.Getenv(k); v != "" {
return v
}
return d
}

func (c *Config) DSN() string {
return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}
