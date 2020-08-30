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
	"github.com/sirupsen/logrus"
)

var (
	data = flag.String("data", "./data.db", "Path for storing data")
	port = flag.String("port", "8080", "Service HTTP port")
)

func main() {
	log.Printf("Starting...")
	logrus.SetLevel(logrus.DebugLevel)

	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		cancel()
	}()
	wg := &sync.WaitGroup{}

	log.Printf("Opening DB...")
	db, err := sql.Open("sqlite3", *data)
	if err != nil {
		log.Fatalf("Failed to create SQLite DB: %s", err)
	}
	defer db.Close()

	log.Printf("Creating controllers...")
	events := make(chan controller.Event, 1)
	firmatasController := controller.NewFirmatas(events, db)
	atoController := controller.NewATO(events, db, firmatasController)
	awcController := controller.NewAWC(events, db, firmatasController)

	log.Printf("Creating API handlers...")
	graphqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db, firmatasController, atoController, awcController),
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlSrv)
	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: http.DefaultServeMux,
	}

	wg.Add(2)
	go atoController.Run(ctx, wg)
	go awcController.Run(ctx, wg)
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	go func() {
		<-ctx.Done()
		srv.Shutdown(ctx)
	}()

	log.Printf("Ready")
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", *port)
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
