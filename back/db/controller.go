package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController() (*Controller, error) {
	db, err := gorm.Open(postgres.Open("host=postgres user=postgres password=postgres dbname=todo port=5432"))
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Todo{}); err != nil {
		return nil, err
	}
	return &Controller{db}, nil
}

func (c *Controller) CreateNewTodo(todo *Todo) (*Todo, error) {
	tx := c.db.Begin()
	if err := tx.Create(todo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return todo, nil
}

func (c *Controller) FetchAllTodos() ([]Todo, error) {
	todos := []Todo{}
	if err := c.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (c *Controller) FetchTodoByID(id int64) (*Todo, error) {
	todo := &Todo{}
	if err := c.db.Where("id = ?", id).Find(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
