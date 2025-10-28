package main

import (
	"log"
	"time"

	"github.com/donnebaldemeca/RESTAPI/internal/db"
	"github.com/donnebaldemeca/RESTAPI/internal/env"
	"github.com/donnebaldemeca/RESTAPI/internal/storage"
)

func main() {
	env.LoadEnv()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

	log.Printf("DB_DSN: %v", cfg.db.dsn)
	log.Printf("DB_MAX_OPEN_CONNS: %d", cfg.db.maxOpenConns)
	log.Printf("DB_MAX_IDLE_CONNS: %d", cfg.db.maxIdleConns)
	log.Printf("DB_MAX_IDLE_TIME: %v", cfg.db.maxIdleTime)

	db, err := db.New(
		cfg.db.dsn,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	storage := storage.NewStorage(db)

	app := &application{
		config:  cfg,
		storage: storage,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
