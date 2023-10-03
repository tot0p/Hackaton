package controller

import "github.com/gin-gonic/gin"

func DataSetByIdAPIController(ctx *gin.Context) {
	id := ctx.Param("id")

	_ = id

	ctx.JSON(200, gin.H{
		"message": "DataSetByIdAPIController",
	})
}
