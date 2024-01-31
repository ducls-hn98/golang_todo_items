package ginitem

import (
	"net/http"
	"strconv"
	"todo_items/common"
	"todo_items/modules/item/biz"
	"todo_items/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		storage := storage.NewSQLStorage(db)
		business := biz.NewDeleteItemBiz(storage)

		if err := business.DeleteItemByID(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("Success"))
	}
}
