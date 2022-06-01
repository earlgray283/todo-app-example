package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"

	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.c.GetAllTodos(ctx)
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
	todo, err := r.c.GetTodoByID(ctx, id2)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todo, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
