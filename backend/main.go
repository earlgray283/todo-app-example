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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	ProjectID = "learning-346605"
)

const (
	defaultPort     = "8080"
	CredentialsName = "credentials.json"
)

func init() {
	if os.Getenv("TODO_PROJECT_ID") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}
	ProjectID = os.Getenv("TODO_PROJECT_ID")
}

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

	r := gin.Default()
	r.Use(cors.Default())
	r.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		h := playground.Handler("GraphQL playground", "/query")
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.Handle(http.MethodPost, "/query", func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
