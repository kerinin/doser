package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	badger "github.com/dgraph-io/badger/v2"
	"github.com/kerinin/doser/service/graph"
	"github.com/kerinin/doser/service/graph/generated"
)

var (
	data = flag.String("data", "", "Path for storing data")
	port = flag.String("port", "8080", "Service HTTP port")
)

func main() {
	var (
		db  *badger.DB
		err error
	)

	if *data == "" {
		db, err = badger.Open(badger.DefaultOptions("").WithInMemory(true))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db, err = badger.Open(badger.DefaultOptions(*data))
		if err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db),
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
