package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"firebase.google.com/go/auth"
	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (r *todoResolver) UserID(ctx context.Context, obj *model.Todo) (string, error) {
	user, _ := ctx.Value(userKey).(*auth.UserRecord)
	if user == nil {
		return "", errors.New("To use this feature, you must login")
	}
	return user.UID, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
