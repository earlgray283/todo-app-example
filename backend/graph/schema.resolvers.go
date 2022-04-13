package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/earlgray283/todo-graphql-firestore/graph/generated"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	newTodo := &model.Todo{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Done:        false,
	}
	if err := r.c.RegistTodo(ctx, newTodo); err != nil {
		return nil, err
	}
	return newTodo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.c.GetAllTodos(ctx)
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	id2, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.c.GetTodoByID(ctx, id2)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
