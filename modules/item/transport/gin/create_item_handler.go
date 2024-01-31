package ginitem

import (
	"net/http"
	"todo_items/common"
	"todo_items/modules/item/biz"
	"todo_items/modules/item/model"
	"todo_items/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
