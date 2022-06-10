package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	userId, _ := ctx.Value(userKey).(*auth.UserRecord)
	if userId == nil {
		return nil, errors.New("userId was nil")
	}
	newTodo := &model.Todo{
		UserID:      userId.UID,
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Done:        false,
	}
	if err := r.fc.RegisterTodo(ctx, newTodo); err != nil {
		log.Println(err)
		return nil, err
	}
	return newTodo, nil
}

func (r *mutationResolver) SessionLogin(ctx context.Context, token string) (*model.SessionToken, error) {
	client, err := r.fb.Auth(ctx)
	if err != nil {
		return nil, err
	}
	expiresIn := time.Hour * 24 * 7
	cookie, err := client.SessionCookie(ctx, token, expiresIn)
	if err != nil {
		return nil, err
	}
	rw, _ := ctx.Value(respWriterKey).(http.ResponseWriter)
	if rw == nil {
		return nil, errors.New("sessionCookie was nil")
	}
	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   false, // because local serve is http

	}
	http.SetCookie(rw, sessionCookie)
	return &model.SessionToken{Token: cookie}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
