package storage

import (
	"context"
	"todo_items/modules/item/model"
)

func (sql *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
