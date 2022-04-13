package firestore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/earlgray283/todo-graphql-firestore/model"
)

func (ctrler *Controller) RegistTodo(ctx context.Context, todo *model.Todo) error {
	imcompleteKey := datastore.IncompleteKey("todos", nil)
	todo.CreatedAt = time.Now()
	key, err := ctrler.c.Put(ctx, imcompleteKey, todo)
	if err != nil {
		return err
	}
	todo.ID = new(string)
	*todo.ID = fmt.Sprint(key.ID)
	return nil
}
