package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *mySqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	deletdeStatus := model.ItemStatusDeleted

	if err := s.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletdeStatus.String(),
		}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
