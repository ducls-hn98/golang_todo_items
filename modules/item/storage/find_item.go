package storage

import (
	"context"
	"todo_items/modules/item/model"
)

func (sql *sqlStorage) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
