package controllers

import (
	"category-service/models"

	"github.com/gin-gonic/gin"
)

func HandleCategoryList(ctx *gin.Context) {
	categories, err := models.GetListCategory()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Category List. Err:" + err.Error(),
		})
		return
	}
	ctx.JSON(200, categories)
}
