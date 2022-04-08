package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/earlgray283/todo-graphql-firestore/firestore"
)

type Resolver struct {
	c *firestore.Controller
}

func NewResolver(controller *firestore.Controller) *Resolver {
	return &Resolver{c: controller}
}
