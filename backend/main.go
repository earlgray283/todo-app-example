package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/earlgray283/todo-graphql-firestore/firestore"
	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"github.com/earlgray283/todo-graphql-firestore/graph/resolver"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const (
	defaultPort = "8080"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file was not found. Use os environment values.")
	}
}

func main() {
	ctx := context.Background()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Println("PORT", os.Getenv("PORT"))
	log.Println("TODO_PROJECT_ID", os.Getenv("TODO_PROJECT_ID"))
	log.Println("TODO_FRONTEND_URL", os.Getenv("TODO_FRONTEND_URL"))

	controller, err := firestore.NewController(ctx, os.Getenv("TODO_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	client, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile("firebase_credentials.json"))
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.NewResolver(controller, client),
			},
		),
	)

	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000", os.Getenv("TODO_FRONTEND_URL")}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig), resolver.MiddlewareSessionCookie(), resolver.MiddlewareAuth(client))
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
