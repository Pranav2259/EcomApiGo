package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	db_dsn := os.Getenv("GOOSE_DBSTRING")

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db:   dbconfig{dsn: db_dsn},
	}

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	app := app{
		config: cfg,
		db:     conn,
	}

	err = app.run(app.mount())
	if err != nil {
		log.Fatalf("Error in runnning server %s", err)
	}
}
