package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}

type createItembusiness struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItembusiness {
	return &createItembusiness{store: store}
}

func (biz *createItembusiness) CreateNewItem(ctx context.Context, data *model.TodoItemCreate) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
