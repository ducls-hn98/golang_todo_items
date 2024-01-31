package ginitem

import (
	"net/http"
	"strconv"
	"todo_items/common"
	"todo_items/modules/item/biz"
	"todo_items/modules/item/model"
	"todo_items/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemByID(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
