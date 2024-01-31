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

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		paging.Process()

		var filter model.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data []model.TodoItem

		store := storage.NewSQLStorage(db)
		business := biz.NewListItemBiz(store)
		data, err := business.ListItem(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
