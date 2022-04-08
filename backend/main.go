package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/earlgray283/todo-graphql-firestore/firestore"
	"github.com/earlgray283/todo-graphql-firestore/graph"
	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"google.golang.org/api/option"
)

const (
	defaultPort     = "8080"
	ProjectID       = "learning-346605"
	CredentialsName = "credentials.json"
)

func main() {
	ctx := context.Background()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	controller, err := firestore.NewController(ctx, ProjectID, option.WithCredentialsFile(CredentialsName))
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: graph.NewResolver(controller),
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
