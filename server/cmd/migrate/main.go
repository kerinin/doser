package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3"

	migrations "github.com/kerinin/doser/service/sql"
)

var (
	data = flag.String("data", "./data.db", "Path for storing data")
)

func main() {
	flag.Parse()

	log.Printf("Migrating %s", *data)

	db, err := sql.Open("sqlite3", *data)
	if err != nil {
		log.Fatalf("Failed to create SQLite DB: %s", err)
	}
	defer db.Close()

	applied, err := migrations.Migrate(db, "sqlite3")
	if err != nil {
		log.Fatalf("Failed to migrate DB: %s", err)
	}
	if applied > 0 {
		log.Printf("Applied %d migrations", applied)
	}

	log.Print("Done")
}
