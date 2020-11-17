package main

import (
	"log"
	"net/http"
	"os"

	"github.com/4saito5/live-streaming-recorder/graph"
	"github.com/4saito5/live-streaming-recorder/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("/data/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	// defer func() {
	// 	if db != nil {
	// 		if err := db.Close(); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }()


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
