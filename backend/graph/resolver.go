package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	firebase "firebase.google.com/go/v4"
	"github.com/earlgray283/todo-graphql-firestore/firestore"
)

type Resolver struct {
	fc *firestore.Controller
	fb *firebase.App
}

func NewResolver(fc *firestore.Controller, fb *firebase.App) *Resolver {
	return &Resolver{fc, fb}
}
