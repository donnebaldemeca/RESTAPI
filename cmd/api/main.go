package main

import (
	"log"
	"os"

	"github.com/donnebaldemeca/RESTAPI/internal/env"
)

func main() {
	env.LoadEnv()

	cfg := config{
		addr: os.Getenv("ADDR"),
	}

	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
