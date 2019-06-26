package controllers

import (
	"product-service/models"

	"github.com/gin-gonic/gin"
)

func HandleProductList(ctx *gin.Context) {
	products, err := models.GetListProduct()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Product List. Err:" + err.Error(),
		})
		return
	}
	ctx.JSON(200, products)
}
