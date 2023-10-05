package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/utils/db/mongodb"
)

// DataSetNamesAPIController handles the GET request to get the list of all datasets names
func DataSetNamesAPIController(ctx *gin.Context) {
	names, err := mongodb.DB.GetAllCollectionsNames(env.Get("DB_MONGODB"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, names)
}
