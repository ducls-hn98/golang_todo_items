package storage

import (
	"context"
	"todo_items/modules/item/model"
)

func (sql *sqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := sql.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{"status": "Deleted"}).Error; err != nil {
		return err
	}

	return nil
}
