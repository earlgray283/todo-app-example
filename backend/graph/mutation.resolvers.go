package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

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
	if err := r.c.RegisterTodo(ctx, newTodo); err != nil {
		log.Println(err)
		return nil, err
	}
	return newTodo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
