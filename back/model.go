package main

import (
	"sync"
	"time"
)

type Todo struct {
	Title       string
	Description string
	createdAt   time.Time
}

type TodoStorage struct {
	sync.Mutex
	Todos []Todo
}

var todoMap = map[string]Todo{
	"0": {
		Title:       "test1",
		Description: "hogehoge",
		createdAt:   time.Now(),
	},
	"1": {
		Title:       "test2",
		Description: "hogehoge",
		createdAt:   time.Now(),
	},
	"2": {
		Title:       "test3",
		Description: "hogehoge",
		createdAt:   time.Now(),
	},
}

var todoStorage TodoStorage
