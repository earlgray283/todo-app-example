package firestore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/earlgray283/todo-graphql-firestore/model"
	"github.com/pkg/errors"
)

const KindTodo = "todos"

func (ctrler *Controller) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	_, err := ctrler.c.GetAll(ctx, datastore.NewQuery(KindTodo), &todos)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return todos, nil
}

func (ctrler *Controller) GetTodoByID(ctx context.Context, id int64) (*model.Todo, error) {
	var todo model.Todo
	if err := ctrler.c.Get(ctx, datastore.IDKey(KindTodo, id, nil), &todo); err != nil {
		return nil, errors.WithStack(err)
	}
	return &todo, nil
}

func (ctrler *Controller) RegisterTodo(ctx context.Context, todo *model.Todo) error {
	todo.CreatedAt = time.Now()
	key := datastore.NameKey(KindTodo, fmt.Sprintf("%v_%v", todo.UserID, todo.CreatedAt.Format(DateFormat)), nil)
	_, err := ctrler.c.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		_, err := tx.Put(key, todo)
		return err
	})
	return err
}
