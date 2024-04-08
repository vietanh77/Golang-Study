package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItembusiness struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItembusiness {

	return &getItembusiness{store: store}
}

func (biz *getItembusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}
	return data, nil
}
