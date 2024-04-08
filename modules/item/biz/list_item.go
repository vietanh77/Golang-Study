package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItembusiness struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItembusiness {

	return &listItembusiness{store: store}
}

func (biz *listItembusiness) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, common.CannotListEmpty(model.EntityName, err)
	}
	return data, nil
}
