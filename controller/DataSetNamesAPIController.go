package controller

import (
	"github.com/gin-gonic/gin"
	"hackaton/utils/db/mongodb"
)

func DataSetNamesAPIController(ctx *gin.Context) {
	names, err := mongodb.DB.GetAllCollectionsNames("DataSet")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, names)
}
