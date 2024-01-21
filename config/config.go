package config

import (
	"flag"
	"os"
)

var (
	Db_conn    string
	Redis_conn string
	Ws_host    string
)

// InitConfig initializes the configuration values.
func InitConfig() {
	// Check if flags are provided
	flag.StringVar(&Db_conn, "db-conn", getEnv("db-conn", "host=localhost port=5432 user=postgres password=123456 sslmode=disable dbname=postgres"), "Database connection string")
	flag.StringVar(&Redis_conn, "redis-conn", getEnv("redis-conn", "localhost:6379"), "Redis connection string")
	flag.StringVar(&Ws_host, "ws-host", getEnv("ws-host", "0.0.0.0:8081"), "WebSocket host")

	flag.Parse()

}
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
