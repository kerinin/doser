package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		cancel()
	}()
	wg := &sync.WaitGroup{}

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

	events := make(chan controller.Event, 1)
	atoControl := controller.NewATOControl(events, db, gomatas)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db),
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	wg.Add(1)
	go atoControl.Run(ctx, wg)
	go log.Fatal(http.ListenAndServe(":"+*port, nil))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	for {
		select {
		case event := <-events:
			// TODO: Handle these more robustly
			log.Printf("Event: %s", event.Message())
		case <-ctx.Done():
			log.Printf("Shutting down...")
			wg.Wait()
			return
		}
	}
}
