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
	_ "net/http/pprof"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kerinin/doser/service/controller"
	"github.com/kerinin/doser/service/graph"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var (
	data    = flag.String("data", "./data.db", "Path for storing data")
	port    = flag.String("port", "8080", "Service HTTP port")
	verbose = flag.Bool("verbose", false, "Verbose logging")
)

func main() {
	flag.Parse()

	log.Printf("Starting...")
	if *verbose {
		log.Printf("Verbose logging enabled")
		logrus.SetLevel(logrus.DebugLevel)
	}

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
	firmatasController := controller.NewFirmatas(db)
	atoController := controller.NewATO(db, firmatasController)
	awcController := controller.NewAWC(db, firmatasController)

	log.Printf("Creating API handlers...")
	graphqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db, firmatasController, atoController, awcController),
	}))

	mux := http.NewServeMux()
	mux.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	mux.HandleFunc("/query", graphqlSrv.ServeHTTP)
	handler := cors.Default().Handler(mux)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: handler,
	}

	wg.Add(2)
	go atoController.Run(ctx, wg)
	go awcController.Run(ctx, wg)
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
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
		case <-ctx.Done():
			log.Printf("Shutting down...")
			wg.Wait()
			return
		}
	}
}
