package model

import (
	"errors"
	"social-todo-list/common"
)

const EntityName = "Item"

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"colum: title"`
	Description string      `json:"description" gorm:"colum: description"`
	Status      *ItemStatus `json:"status" gorm:"colum: status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreate struct {
	Id          int         `json:"-" gorm:"colum:id" `
	Title       string      `json:"title" gorm:"colum:title"`
	Description string      `json:"description" gorm:"colum:description"`
	Status      *ItemStatus `json:"status" gorm:"colum:status"`
}

func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"colum:title"`
	Description *string     `json:"description" gorm:"colum:description"`
	Status      *ItemStatus `json:"status" gorm:"colum:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
