package db

import (
	"time"
)

type Todo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duedate     time.Time `json:"duedate"`

	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
