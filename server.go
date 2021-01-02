package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/subh1994/graphql-go-server/graph"
	"github.com/subh1994/graphql-go-server/graph/generated"
	"github.com/subh1994/graphql-go-server/internal/pkg/db/postgres"
)

const defaultPort = "5000"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading env", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	postgres.InitDbPool()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
