package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kerinin/doser/service/controller"
	"github.com/kerinin/doser/service/graph"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/kerinin/doser/service/models"
	"github.com/kerinin/gomata"
)

var (
	data = flag.String("data", "./data.db", "Path for storing data")
	port = flag.String("port", "8080", "Service HTTP port")
)

func main() {
	// TODO: Sig hanling
	ctx := context.Background()

	db, err := sql.Open("sqlite3", *data)
	if err != nil {
		log.Fatalf("Failed to create SQLite DB: %s", err)
	}
	defer db.Close()

	firmatas, err := models.Firmatas().All(ctx, db)
	if err != nil {
		log.Fatalf("Failed to get list of firmatas: %s", err)
	}
	gomatas := make(map[string]*gomata.Firmata, len(firmatas))
	for _, firmata := range firmatas {
		gomatas[firmata.ID] = gomata.New()
	}

	// TODO: Handle errors from the controller
	atoControl := controller.NewATOControl(db, gomatas)
	go atoControl.Run(ctx)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db),
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
