package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (ctrler *Controller) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	keys, err := ctrler.c.GetAll(ctx, datastore.NewQuery("todos"), &todos)
	if err != nil {
		return nil, err
	}
	for i := range keys {
		todos[i].ID = new(string)
		*todos[i].ID = fmt.Sprint(keys[i].ID)
	}
	return todos, nil
}

func (ctrler *Controller) GetTodoByID(ctx context.Context, id int64) (*model.Todo, error) {
	var todo model.Todo
	if err := ctrler.c.Get(ctx, datastore.IDKey("todos", id, nil), &todo); err != nil {
		return nil, err
	}
	todo.ID = new(string)
	*todo.ID = fmt.Sprint(id)
	return &todo, nil
}
