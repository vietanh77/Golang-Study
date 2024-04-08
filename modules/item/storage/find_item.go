package storage

import (
	"context"
	"gorm.io/gorm"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *mySqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNoFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
