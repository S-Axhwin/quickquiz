package conn

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {

	dsn := os.Getenv("DATABASE_URL")
	tokenKey := os.Getenv("JWT_SECRET")
	if dsn == "" || tokenKey == "" {
		log.Fatal("DATABASE_URL or JWT_SECRET not set")
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	// ---- REALISTIC DEFAULTS (important) ----
	config.MaxConns = 20 // tune based on traffic
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	// ALWAYS PING
	if err := Pool.Ping(ctx); err != nil {
		log.Fatalf("database unreachable: %v", err)
	}

	log.Println("Postgres connected with pgxpool")
}
