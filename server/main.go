package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kerinin/doser/service/graph"
	"github.com/kerinin/doser/service/graph/generated"
	migrations "github.com/kerinin/doser/service/sql"
)

var (
	data = flag.String("data", "./data.db", "Path for storing data")
	port = flag.String("port", "8080", "Service HTTP port")
)

func main() {
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db),
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
