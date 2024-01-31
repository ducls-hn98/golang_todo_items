package model

import (
	"errors"
	"time"
	"todo_items/common"
)

var (
	ErrTitleIsBlank = errors.New("title can not be blank")
	ErrItemDeleted = errors.New("this item is already deleted")
)

type TodoItem struct {
	//Embedded struct
	common.SQLModel
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	ID          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:description;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string    `json:"title" gorm:"column:title"`
	Description *string    `json:"description" gorm:"column:description"`
	Status      *string    `json:"status" gorm:"column:status"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
