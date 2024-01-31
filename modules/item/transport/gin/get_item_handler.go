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

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewGetItemBiz(store)
		data, err := business.GetItemByID(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
