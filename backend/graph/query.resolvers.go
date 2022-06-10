package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"strconv"

	"firebase.google.com/go/v4/auth"
	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	user, _ := ctx.Value(userKey).(*auth.UserRecord)
	if user == nil {
		return nil, errors.New("user must not be nil")
	}
	todos, err := r.fc.GetAllTodosByUserID(ctx, user.UID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	id2, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	todo, err := r.fc.GetTodoByID(ctx, id2)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todo, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
