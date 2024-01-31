package storage

import (
	"context"
	"todo_items/modules/item/model"
)

func (sql *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
